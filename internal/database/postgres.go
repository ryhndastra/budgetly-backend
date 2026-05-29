package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Connect(connString string) (*pgx.Conn, error) {
	return pgx.Connect(
		context.Background(),
		connString,
	)
}