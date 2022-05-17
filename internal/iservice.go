package internal

import "fabric-voter/internal/models"

type Service interface {
	CreateThread(params *models.ThreadParams) (string, error)
	GetThread(threadID string) (*models.Thread, error)
	CreateVote(threadID string) (*models.Vote, error)
	UseVote(vote *models.Vote) error
	EndThread(threadID string) error

	CreateAnonThread(params *models.ThreadParams) (string, error)
	GetAnonThread(threadID string) (*models.AnonThread, error)
	UseAnonVote(vote *models.AnonVote) error
	EndAnonThread(data *models.EndAnonData) error
}