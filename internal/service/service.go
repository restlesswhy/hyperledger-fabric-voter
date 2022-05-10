package service

import (
	"fabric-voter/internal"
	"fabric-voter/internal/models"
)

type service struct {
	repo   internal.Repository
	ledger internal.Ledger
}

func NewService(repo internal.Repository, ledger internal.Ledger) internal.Service {
	return &service{
		repo:   repo,
		ledger: ledger,
	}
}

func (s *service) CreateThread(params *models.ThreadParams) error {
	return nil
}

func (s *service) GetThread(threadID string) (*models.Thread, error) {
	return nil, nil
}

func (s *service) CreateVote(threadID string) (string, error) {
	return "", nil
}

func (s *service) UseVote(vote *models.Vote) error {
	return nil
}

func (s *service) EndThread(threadID string) error {
	return nil
}
