package config

import (
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Ledger   LedgerConfig
}

type PostgresConfig struct {
	Username string
	Password string
	Hostname string
	Port     int
	DBName   string
}

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type LedgerConfig struct {
	MspID         string
	CertPath      string
	KeyPath       string
	TlsCertPath   string
	PeerEndpoint  string
	GatewayPeer   string
	ChannelName   string
	ChaincodeName string
}

func Load(filename string) (*Config, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath("./config")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		logrus.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
