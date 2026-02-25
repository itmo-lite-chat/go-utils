package pgdb

import (
	"context"
	"database/sql"
	"embed"

	"github.com/itmo-lite-chat/go-utils/logger"
	"github.com/pkg/errors"
	"github.com/pressly/goose/v3"
)

func ApplyMigrations(ctx context.Context, db *sql.DB, embedMigrations embed.FS) error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return errors.Wrap(err, "unable to select dialect of migrations")
	}

	logger.Info(ctx, "running migrations")
	if err := goose.Up(db, "migrations"); err != nil {
		return errors.Wrap(err, "unable to apply migrations")
	}
	logger.Info(ctx, "migrations applied successfully")

	return nil
}
