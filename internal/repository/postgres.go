package repository

import (
	"context"
	"fabric-voter/internal"
	"fabric-voter/internal/models"
	"fmt"
	"strings"

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
	options := strings.Join(params.Options, `", "`)
	options = fmt.Sprintf(`{"%s"}`, options)

	q := `INSERT INTO threads (
			thread_id,
			category,
			theme,
			description,
			options
		) VALUES (
			$1, $2, $3, $4, $5
		);`
	r.pool.Exec(context.Background(), q, params.ID, params.Category, params.Theme, params.Description, options)

	return nil
}

func (r *repo) GetThread(threadID string) (*models.Thread, error) {
	q := `SELECT `

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
