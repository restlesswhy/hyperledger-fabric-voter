package internal

import "fabric-voter/internal/models"

type Repository interface {
	CreateThread(threadID string, thread *models.Thread) error
	GetThread(threadID string) (*models.Thread, error)
	CreateVote(threadID string) (string, error)
	UseVote(vote *models.Vote) error
	EndThread(threadID string) error
}