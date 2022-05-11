package repository

import (
	"context"
	"encoding/json"
	"fabric-voter/internal"
	"fabric-voter/internal/models"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type repo struct {
	pool *pgxpool.Pool
}

func NewRepo(pool *pgxpool.Pool) internal.Repository {
	return &repo{
		pool: pool,
	}
}

func (r *repo) CreateThread(threadID string, thread *models.Thread) error {
	t, err := json.Marshal(thread)
	if err != nil {
		return errors.Wrap(err, "err marshal thread")
	}
	
	q := `INSERT INTO threads (thread_id, thread) VALUES ($1, $2);`
	f := pgtype.JSONB{}
	err = f.Set(t)
	if err != nil {
		return errors.Wrap(err, "err jsonb bytes")
	}
	// q := `INSERT INTO threads
	// 		SELECT thread_id, category, theme, description, creator, options, win_option, status
	// 		FROM json_populate_record (NULL::threads, $1);`

	_, err = r.pool.Exec(context.Background(), q, threadID, f)
	if err != nil {
		return errors.Wrap(err, "err exec thread")
	}

	return nil
}

func (r *repo) GetThread(threadID string) (*models.Thread, error) {
	q := `SELECT thread
			FROM threads
			WHERE thread_id = $1;`
	
	j := pgtype.JSONB{}
	err := r.pool.QueryRow(context.Background(), q, threadID).Scan(&j)
	if err != nil {
		return nil, err
	}

	res := &models.Thread{}
	err = json.Unmarshal(j.Bytes, &res)
	if err != nil {
		return nil, errors.Wrap(err, "err unmarshal thread")
	}
	
	return res, nil
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
