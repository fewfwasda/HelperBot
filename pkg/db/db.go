package db

import (
	"HelperBot/internal/config"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Database struct {
	Pool *pgxpool.Pool
	DB   *sql.DB
}

func NewDatabase(ctx context.Context, cfg config.DBConfig) (*Database, error) {
	poolConfig, err := pgxpool.ParseConfig(cfg.DBUrl)
	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = cfg.MaxConns
	poolConfig.MinConns = cfg.MinConns
	poolConfig.MaxConnIdleTime = cfg.MaxConnsIdle

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("error while creating pool: %w", err)
	}

	ctxPing, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctxPing); err != nil {
		pool.Close()
		return nil, fmt.Errorf("err while connect to DB: %w", err)
	}

	db, err := sql.Open("pgx", cfg.DBUrl)
	if err != nil {
		db.Close()
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return &Database{Pool: pool, DB: db}, nil
}
