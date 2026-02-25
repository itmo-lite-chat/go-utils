package pgdb

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"time"
)

func GetConnectionString(username, password, host string, port int, name string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable&search_path=public",
		username, password, host, port, name,
	)
}

func Connect(ctx context.Context, connString string, ttl int, maxOpenConnections, maxIdleConnections int) (*sql.DB, error) {
	conn, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, errors.Wrap(err, "open db error")
	}

	if err := conn.PingContext(ctx); err != nil {
		return nil, errors.Wrap(err, "ping db error")
	}

	conn.SetConnMaxLifetime(time.Second * time.Duration(ttl))
	conn.SetMaxOpenConns(maxOpenConnections)
	conn.SetMaxIdleConns(maxIdleConnections)

	return conn, nil
}
