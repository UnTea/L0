package helpers

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

// NewPostgres - create new postgres connection
func NewPostgres(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	connection, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	err = connection.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return connection, nil
}
