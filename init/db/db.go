package db

import (
	"HelperBot/internal/config"
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDatabase(cfg config.DBConfig) (*sql.DB, *pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), cfg.DBUrl)
	if err != nil {
		return nil, nil, err
	}

	db, err := sql.Open("pgx", cfg.DBUrl)
	if err != nil {
		pool.Close()
		return nil, nil, err
	}

	if err := db.Ping(); err != nil {
		pool.Close()
		return nil, nil, err
	}

	return db, pool, nil
}
