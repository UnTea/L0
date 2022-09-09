package helpers

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgres(ctx context.Context, connectionString string) (*pgxpool.Pool, error) {
	connection, err := pgxpool.Connect(ctx, connectionString)
	if err != nil {
		return nil, err
	}

	err = connection.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return connection, nil
}
