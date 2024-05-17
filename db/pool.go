package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDbPool(ctx context.Context, dbConfig DbConfig) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), dbConfig.toDatabaseUrl())
	return pool, err
}
