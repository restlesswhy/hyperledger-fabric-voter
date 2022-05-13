package synchronizer

import (
	"bytes"
	"context"
	"encoding/json"
	"fabric-voter/config"
	"fabric-voter/internal"
	"fmt"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

const (
	EVENT_CreateThread = "CreateThread"
	EVENT_UseVote      = "UseVote"
	EVENT_EndThread    = "EndThread"
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
		panic(fmt.Errorf("failed to start chaincode event listening: %w", err))
	}
	
	for event := range events {
		asset := formatJSON(event.Payload)
		switch event.ChaincodeName {
		case EVENT_CreateThread:
			fmt.Printf("\n<-- Chaincode event received: %s - %s\n", event.EventName, asset)
		case EVENT_UseVote:
			fmt.Printf("\n<-- Chaincode event received: %s - %s\n", event.EventName, asset)
		case EVENT_EndThread:
			fmt.Printf("\n<-- Chaincode event received: %s - %s\n", event.EventName, asset)
		}
	}
}

func formatJSON(data []byte) string {
	var result bytes.Buffer
	if err := json.Indent(&result, data, "", "  "); err != nil {
		panic(fmt.Errorf("failed to parse JSON: %w", err))
	}
	return result.String()
}
