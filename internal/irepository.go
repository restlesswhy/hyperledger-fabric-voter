package internal

import "fabric-voter/internal/models"

type Repository interface {
	CreateThread(threadID string, thread []byte) error
	UpdateThread(threadID string, thread []byte) error
	GetThread(threadID string) (*models.Thread, error)

	GetAnonThread(threadID string) (*models.AnonThread, error)
}
