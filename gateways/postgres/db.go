package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/jackc/pgx/v5/pgxpool"
)

type db interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type scan func(destination ...any) error

func NewConnectionPool(ctx context.Context, addr string, minConn, maxConn int32) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(addr)
	if err != nil {
		return nil, err
	}

	cfg.MinConns = minConn
	cfg.MaxConns = maxConn

	pgxConnPoll, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return pgxConnPoll, nil
}
