package db

import (
	"context"
	"database/sql"
	"fmt"
)

func ConnectDB(ctx context.Context, dsn string) (*sql.DB, error) {
	conn, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, fmt.Errorf("opening db connection: %w", err)
	}

	if err := conn.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("pinging db: %w", err)
	}

	return conn, nil
}
