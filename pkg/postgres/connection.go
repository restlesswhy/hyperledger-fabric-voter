package postgres

import (
	"fabric-voter/config"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Connect(cfg *config.Config) (*pgxpool.Pool, error) {
	databaseUrl := fmt.Sprintf(`postgres://%s:%s@%s:%d/%s`,
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.Hostname,
		cfg.Postgres.Port,
		"postgres")

	dbpool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		return nil, err
	}

	if err = dbpool.Ping(context.Background()); err != nil {
		return nil, err
	}

	var exists bool
	if err := dbpool.QueryRow(context.Background(), `SELECT COUNT(*)>0 AS db_exists FROM pg_database WHERE datname = $1;`, cfg.Postgres.DBName).Scan(&exists); err != nil {
		return nil, fmt.Errorf("cannot check if database exists: %v", err)
	}

	if !exists {
		if err := CreateDB(dbpool, cfg.Postgres.DBName); err != nil {
			return nil, fmt.Errorf("cannot init database: %v", err)
		}
	}
	dbpool.Close()

	databaseUrl = fmt.Sprintf(`postgres://%s:%s@%s:%d/%s`,
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.Hostname,
		cfg.Postgres.Port,
		cfg.Postgres.DBName)

	dbpool, err = pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		return nil, err
	}

	err = dbpool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	err = initDB(dbpool)
	if err != nil {
		return nil, err
	}

	return dbpool, nil
}

func CreateDB(pool *pgxpool.Pool, dbName string) error {
	query := `CREATE DATABASE ` + dbName + ";"
	_, err := pool.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}

func initDB(pool *pgxpool.Pool) error {
	query := `
		CREATE TABLE IF NOT EXISTS buttons (
			button_id integer NOT NULL PRIMARY KEY,
			message_id integer NOT NULL UNIQUE,
			amount integer NOT NULL,
			firstname varchar(250) NOT NULL,
			lastname varchar(250) NOT NULL,
			username varchar(250) NOT NULL
		);

		CREATE TABLE IF NOT EXISTS transactions (
			tx_id SERIAL PRIMARY KEY,
			button_id integer NOT NULL,
			amount integer NOT NULL,
			category varchar(200) NOT NULL,
			time timestamp NOT NULL,
			CONSTRAINT fk_button_id
				FOREIGN KEY(button_id)
				REFERENCES buttons(button_id) 
		);
	`

	_, err := pool.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}
