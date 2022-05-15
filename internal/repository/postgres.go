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

func (r *repo) CreateThread(threadID string, thread []byte) error {
	q := `INSERT INTO threads
			(thread_id, thread)
			VALUES ($1, $2);`
	jb := pgtype.JSONB{}
	err := jb.Set(thread)
	if err != nil {
		return errors.Wrap(err, "jb.Set()")
	}

	_, err = r.pool.Exec(context.Background(), q, threadID, jb)
	if err != nil {
		return errors.Wrap(err, "r.pool.Exec()")
	}

	return nil
}

func (r *repo) GetThread(threadID string) (*models.Thread, error) {
	q := `SELECT thread
			FROM threads
			WHERE thread_id = $1;`

	jb := pgtype.JSONB{}
	err := r.pool.QueryRow(context.Background(), q, threadID).Scan(&jb)
	if err != nil {
		return nil, errors.Wrap(err, "r.pool.QueryRow()")
	}

	res := &models.Thread{}
	err = json.Unmarshal(jb.Bytes, &res)
	if err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal()")
	}

	return res, nil
}

func (r *repo) UpdateThread(threadID string, thread []byte) error {
	q := `UPDATE threads
			SET thread = $1
			WHERE thread_id = $2;`

	jb := pgtype.JSONB{}
	err := jb.Set(thread)
	if err != nil {
		return errors.Wrap(err, "jb.Set()")
	}

	_, err = r.pool.Exec(context.Background(), q, jb, threadID)
	if err != nil {
		return errors.Wrap(err, "r.pool.Exec()")
	}

	return nil
}