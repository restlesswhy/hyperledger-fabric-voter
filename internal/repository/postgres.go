package repository

import (
	"fabric-voter/internal"
	"fabric-voter/internal/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type repo struct {
	pool *pgxpool.Pool
}

func NewRepo(pool *pgxpool.Pool) internal.Repository {
	return &repo{
		pool: pool,
	}
}

func (r *repo) CreateThread(params *models.ThreadParams) error {

	return nil
}

func (r *repo) GetThread(threadID string) (*models.Thread, error) {

	return nil, nil
}

func (r *repo) CreateVote(threadID string) (string, error) {

	return "", nil
}

func (r *repo) UseVote(vote *models.Vote) error {

	return nil
}

func (r *repo) EndThread(threadID string) error {
	return nil
}
