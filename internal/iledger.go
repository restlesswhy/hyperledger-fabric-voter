package internal

import "fabric-voter/internal/models"

type Ledger interface {
	CreateThread(params *models.ThreadParams) error
	GetThread(threadID string) (*models.Thread, error)
	CreateVote(threadID string, userID string) (*models.Vote, error)
	UseVote(vote *models.Vote) error
	EndThread(threadID string) error

	CreateAnonThread(params *models.ThreadParams) error
	GetAnonThread(threadID string) (*models.AnonThread, error)
	UseAnonVote(vote *models.AnonVote) error
	EndAnonThread(data *models.EndAnonData) error	
}