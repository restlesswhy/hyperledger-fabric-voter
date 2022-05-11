package server

import (
	"crypto/x509"
	"fabric-voter/config"
	"fabric-voter/internal/handler"
	"fabric-voter/internal/ledger"
	"fabric-voter/internal/repository"
	"fabric-voter/internal/service"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"path"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	maxHeaderBytes = 1 << 20
)

type Server struct {
	cfg  *config.Config
	pool *pgxpool.Pool
}

func NewServer(cfg *config.Config, pool *pgxpool.Pool) *Server {
	return &Server{
		cfg:  cfg,
		pool: pool,
	}
}

func (s *Server) Run() error {
	logrus.SetLevel(logrus.DebugLevel)
	clientConnection := newGrpcConnection(s.cfg)
	defer clientConnection.Close()

	id := newIdentity(s.cfg)
	sign := newSign(s.cfg)

	// Create a Gateway connection for a specific client identity
	gateway, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		return err
	}
	defer gateway.Close()

	network := gateway.GetNetwork(s.cfg.Ledger.ChannelName)
	contract := network.GetContract(s.cfg.Ledger.ChaincodeName)

	repo := repository.NewRepo(s.pool)
	ledger := ledger.NewLedger(contract)

	// params := &models.ThreadParams{
	// 	ID: "thread2",
	// 	Theme: "Who wand to eat?",
	// 	Options: []string{"a", "b", "c"},
	// }

	// err = ledger.CreateThread(params)
	// if err != nil {
	// 	logrus.Fatal(err)
	// }
	// thr, err := ledger.GetThread("thread2")
	// if err != nil {
	// 	logrus.Fatal(err)
	// }
	// fmt.Println(*thr)
	// tx, err := ledger.CreateVote("thread2")
	// if err != nil {
	// 	logrus.Fatal(err)
	// }
	// fmt.Println(tx)
	// vote := &models.Vote{
	// 	ThreadID: "thread2",
	// 	VoteID: tx,
	// 	Option: "a",
	// }
	// err = ledger.UseVote(vote)
	// if err != nil {
	// 	logrus.Fatal(err)
	// }
	// err = ledger.EndThread("thread2")
	// if err != nil {
	// 	logrus.Fatal(err)
	// }
	// thr, err = ledger.GetThread("thread2")
	// if err != nil {
	// 	logrus.Fatal(err)
	// }
	// fmt.Println(*thr)
	// ledger.InitLedger()
	// ledger.ReadAssetByID()

	service := service.NewService(repo, ledger)
	handler := handler.NewHandler(service)

	server := &http.Server{
		Addr:           s.cfg.Server.Port,
		Handler:        handler.MapRoutes(),
		ReadTimeout:    time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout:   time.Second * s.cfg.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logrus.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	return nil
}

func newGrpcConnection(cfg *config.Config) *grpc.ClientConn {
	certificate, err := loadCertificate(cfg.Ledger.TlsCertPath)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, cfg.Ledger.GatewayPeer)

	connection, err := grpc.Dial(cfg.Ledger.PeerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}

// newIdentity creates a client identity for this Gateway connection using an X.509 certificate.
func newIdentity(cfg *config.Config) *identity.X509Identity {
	certificate, err := loadCertificate(cfg.Ledger.CertPath)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(cfg.Ledger.MspID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}

// newSign creates a function that generates a digital signature from a message digest using a private key.
func newSign(cfg *config.Config) identity.Sign {
	files, err := ioutil.ReadDir(cfg.Ledger.KeyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key directory: %w", err))
	}
	privateKeyPEM, err := ioutil.ReadFile(path.Join(cfg.Ledger.KeyPath, files[0].Name()))

	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}
