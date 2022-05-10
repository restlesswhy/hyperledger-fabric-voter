package ledger

import (
	"context"
	"encoding/json"
	"errors"
	"fabric-voter/internal"
	"fabric-voter/internal/models"
	"fmt"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	gwproto "github.com/hyperledger/fabric-protos-go/gateway"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

type Ledger interface {
}

type ledger struct {
	client *client.Contract
}

func NewLedger(client *client.Contract) internal.Ledger {
	return &ledger{
		client: client,
	}
}

func (l *ledger) CreateThread(params *models.ThreadParams) error {
	logrus.Debug("Start creating thread...")

	args := make([]string, 0)
	args = append(args, params.ID, params.Category, params.Theme, params.Description)
	args = append(args, params.Options...)

	_, err := l.client.Submit("CreateThread", client.WithArguments(args...))
	if err != nil {
		// return errorHandling(err)
		switch err := err.(type) {
		case *client.EndorseError:
			fmt.Printf("endorse error with gRPC status %v: %s\n", status.Code(err), err)
		case *client.SubmitError:
			return fmt.Errorf("submit error with gRPC status %v: %s", status.Code(err), err)
		case *client.CommitStatusError:
			if errors.Is(err, context.DeadlineExceeded) {
				return fmt.Errorf("timeout waiting for transaction %s commit status: %s", err.TransactionID, err)
			} else {
				return fmt.Errorf("error obtaining commit status with gRPC status %v: %s", status.Code(err), err)
			}
		case *client.CommitError:
			return fmt.Errorf("transaction %s failed to commit with status %d: %s", err.TransactionID, int32(err.Code), err)
		}
		/*
		 Any error that originates from a peer or orderer node external to the gateway will have its details
		 embedded within the gRPC status error. The following code shows how to extract that.
		*/
		statusErr := status.Convert(err)
		for _, detail := range statusErr.Details() {
			errDetail := detail.(*gwproto.ErrorDetail)
			fmt.Printf("error from endpoint: %s, mspId: %s, message: %s\n", errDetail.Address, errDetail.MspId, errDetail.Message)
		}
	}

	logrus.Debug("Thread successfuly created!")
	return nil
}

func (l *ledger) CreateVote(threadID string) (string, error) {
	logrus.Debug("Start creating vote...")
	
	txid, err := l.client.SubmitTransaction("CreateVote", threadID)
	if err != nil {
		// return "", errorHandling(err)
		// return errorHandling(err)
		switch err := err.(type) {
		case *client.EndorseError:
			fmt.Printf("endorse error with gRPC status %v: %s\n", status.Code(err), err)
		case *client.SubmitError:
			// return fmt.Errorf("submit error with gRPC status %v: %s", status.Code(err), err)
		case *client.CommitStatusError:
			if errors.Is(err, context.DeadlineExceeded) {
				// return fmt.Errorf("timeout waiting for transaction %s commit status: %s", err.TransactionID, err)
			} else {
				// return fmt.Errorf("error obtaining commit status with gRPC status %v: %s", status.Code(err), err)
			}
		case *client.CommitError:
			// return fmt.Errorf("transaction %s failed to commit with status %d: %s", err.TransactionID, int32(err.Code), err)
		}
		/*
		 Any error that originates from a peer or orderer node external to the gateway will have its details
		 embedded within the gRPC status error. The following code shows how to extract that.
		*/
		statusErr := status.Convert(err)
		for _, detail := range statusErr.Details() {
			errDetail := detail.(*gwproto.ErrorDetail)
			fmt.Printf("error from endpoint: %s, mspId: %s, message: %s\n", errDetail.Address, errDetail.MspId, errDetail.Message)
		}
	}

	logrus.Debug("Vote successfuly created!")
	return string(txid), nil
}

func (l *ledger) UseVote(vote *models.Vote) error {
	logrus.Debug("Start using anon vote...")

	_, err := l.client.SubmitTransaction("UseVote", vote.ThreadID, vote.VoteID, vote.Option)
	if err != nil {
		// return errorHandling(err)
		switch err := err.(type) {
		case *client.EndorseError:
			fmt.Printf("endorse error with gRPC status %v: %s\n", status.Code(err), err)
		case *client.SubmitError:
			return fmt.Errorf("submit error with gRPC status %v: %s", status.Code(err), err)
		case *client.CommitStatusError:
			if errors.Is(err, context.DeadlineExceeded) {
				return fmt.Errorf("timeout waiting for transaction %s commit status: %s", err.TransactionID, err)
			} else {
				return fmt.Errorf("error obtaining commit status with gRPC status %v: %s", status.Code(err), err)
			}
		case *client.CommitError:
			return fmt.Errorf("transaction %s failed to commit with status %d: %s", err.TransactionID, int32(err.Code), err)
		}
		/*
		 Any error that originates from a peer or orderer node external to the gateway will have its details
		 embedded within the gRPC status error. The following code shows how to extract that.
		*/
		statusErr := status.Convert(err)
		for _, detail := range statusErr.Details() {
			errDetail := detail.(*gwproto.ErrorDetail)
			fmt.Printf("error from endpoint: %s, mspId: %s, message: %s\n", errDetail.Address, errDetail.MspId, errDetail.Message)
		}
	}

	logrus.Debug("Anon vote used!")
	return nil
}

func (l *ledger) EndThread(threadID string) error {
	logrus.Debug("Start closing thread...")

	_, err := l.client.SubmitTransaction("EndThread", threadID)
	if err != nil {
		// return errorHandling(err)
		switch err := err.(type) {
		case *client.EndorseError:
			fmt.Printf("endorse error with gRPC status %v: %s\n", status.Code(err), err)
		case *client.SubmitError:
			// return fmt.Errorf("submit error with gRPC status %v: %s", status.Code(err), err)
		case *client.CommitStatusError:
			if errors.Is(err, context.DeadlineExceeded) {
				// return fmt.Errorf("timeout waiting for transaction %s commit status: %s", err.TransactionID, err)
			} else {
				// return fmt.Errorf("error obtaining commit status with gRPC status %v: %s", status.Code(err), err)
			}
		case *client.CommitError:
			return fmt.Errorf("transaction %s failed to commit with status %d: %s", err.TransactionID, int32(err.Code), err)
		}
		/*
		 Any error that originates from a peer or orderer node external to the gateway will have its details
		 embedded within the gRPC status error. The following code shows how to extract that.
		*/
		statusErr := status.Convert(err)
		for _, detail := range statusErr.Details() {
			errDetail := detail.(*gwproto.ErrorDetail)
			fmt.Printf("error from endpoint: %s, mspId: %s, message: %s\n", errDetail.Address, errDetail.MspId, errDetail.Message)
		}
	}

	logrus.Debug("Thread closed!")
	return nil
}

func (l *ledger) GetThread(threadID string) (*models.Thread, error) {
	logrus.Debug("Start quering thread...")

	res, err := l.client.SubmitTransaction("QueryThread", threadID)
	if err != nil {
		return nil, errorHandling(err)
	}

	thread := &models.Thread{}
	err = json.Unmarshal(res, thread)
	if err != nil {
		return nil, err
	}

	logrus.Debug("Query finded!")
	return thread, nil
}

// Submit transaction, passing in the wrong number of arguments ,expected to throw an error containing details of any error responses from the smart contract.
func errorHandling(err error) error {

	switch err := err.(type) {
	case *client.EndorseError:
		return fmt.Errorf("endorse error with gRPC status %v: %s", status.Code(err), err)
	case *client.SubmitError:
		return fmt.Errorf("submit error with gRPC status %v: %s", status.Code(err), err)
	case *client.CommitStatusError:
		if errors.Is(err, context.DeadlineExceeded) {
			return fmt.Errorf("timeout waiting for transaction %s commit status: %s", err.TransactionID, err)
		} else {
			return fmt.Errorf("error obtaining commit status with gRPC status %v: %s", status.Code(err), err)
		}
	case *client.CommitError:
		return fmt.Errorf("transaction %s failed to commit with status %d: %s", err.TransactionID, int32(err.Code), err)
	}
	/*
	 Any error that originates from a peer or orderer node external to the gateway will have its details
	 embedded within the gRPC status error. The following code shows how to extract that.
	*/
	statusErr := status.Convert(err)
	for _, detail := range statusErr.Details() {
		errDetail := detail.(*gwproto.ErrorDetail)
		return fmt.Errorf("error from endpoint: %s, mspId: %s, message: %s", errDetail.Address, errDetail.MspId, errDetail.Message)
	}

	return err
}
