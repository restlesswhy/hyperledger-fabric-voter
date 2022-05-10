package main

import (
	"os"

	"fabric-voter/config"
	"fabric-voter/internal/server"
	"fabric-voter/pkg/postgres"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err, "error loading env")
	}
	
	cfg, err := config.Load(os.Getenv("CONFIG"))
	if err != nil {
		logrus.Fatal(err, "error loading config")
	}

	pool, err := postgres.Connect(cfg)
	if err != nil {
		logrus.Fatal(err, "error connect to postgres")
	}

	srv := server.NewServer(cfg, pool)

	if err := srv.Run(); err != nil {
		logrus.Fatal(err)
	}
}