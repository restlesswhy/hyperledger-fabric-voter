package synchronizer

import (
	"context"
	"fabric-voter/config"
	"fabric-voter/internal"
	"strings"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/sirupsen/logrus"
)

const (
	EVENT_CreateThread     = "CreateThread"
	EVENT_UseVote          = "UseVote"
	EVENT_EndThread        = "EndThread"
	EVENT_CreateAnonThread = "CreateAnonThread"
	EVENT_UseAnonVote      = "UseAnonVote"
	EVENT_EndAnonThread    = "EndAnonThread"
)

type Synchronizer interface {
	Run(ctx context.Context)
}

type synchronizer struct {
	hf   *client.Network
	repo internal.Repository
	cfg  *config.Config
}

func NewSynchro(hf *client.Network, repo internal.Repository, cfg *config.Config) Synchronizer {
	return &synchronizer{
		hf:   hf,
		repo: repo,
		cfg:  cfg,
	}
}

func (s *synchronizer) Run(ctx context.Context) {
	events, err := s.hf.ChaincodeEvents(ctx, s.cfg.Ledger.ChaincodeName)
	if err != nil {
		logrus.Fatalf("failed to start chaincode event listening: %w", err)
	}

	logrus.Info("start listening events")
	for event := range events {
		spl := strings.Split(event.EventName, " ")
		name, id := spl[0], spl[1]

		switch name {
		case EVENT_CreateThread:
			if err := s.repo.CreateThread(id, event.Payload); err != nil {
				logrus.Warnf("%s failed add into postgres", id)
			}
		case EVENT_UseVote:
			if err := s.repo.UpdateThread(id, event.Payload); err != nil {
				logrus.Warnf("%s failed update in postgres", id)
			}
		case EVENT_EndThread:
			if err := s.repo.UpdateThread(id, event.Payload); err != nil {
				logrus.Warnf("%s failed update in postgres", id)
			}
		case EVENT_CreateAnonThread:
			if err := s.repo.CreateThread(id, event.Payload); err != nil {
				logrus.Warnf("%s failed add into postgres", id)
			}
		case EVENT_UseAnonVote:
			if err := s.repo.UpdateThread(id, event.Payload); err != nil {
				logrus.Warnf("%s failed update in postgres", id)
			}
		case EVENT_EndAnonThread:
			if err := s.repo.UpdateThread(id, event.Payload); err != nil {
				logrus.Warnf("%s failed update in postgres", id)
			}
		}
	}
	logrus.Info("end listening events")
}
