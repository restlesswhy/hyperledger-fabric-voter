package internal

import "fabric-voter/internal/models"

type Ledger interface {
	CreateThread(params *models.ThreadParams) error
	GetThread(threadID string) (*models.Thread, error)
	CreateVote(threadID string) (string, error)
	UseVote(vote *models.Vote) error
	EndThread(threadID string) error
}