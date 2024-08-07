package plain

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgxpool"
)

var MigrationFS fs.FS

// ConnectPgx connects to the database.
//
// Database connection information should be specified by setting standard
// PostgreSQL environment variables, e.g. PGHOST, PGUSER, PGPASSWORD,
// PGDATABASE.
//
// Blocks until a connection has been achieved to prevent application thrashing.
func ConnectPgx(ctx context.Context) (pool *pgxpool.Pool) {
	var err error
	for {
		pool, err = pgxpool.New(ctx, "")
		if err != nil {
			slog.Info("failed to connect to db", "err", err)
			time.Sleep(time.Second)
			continue
		}
		break
	}
	for {
		if err := migrateUp(); err != nil {
			slog.Info("failed to migrate db", "err", err)
			time.Sleep(10 * time.Second)
			continue
		}
		break
	}
	slog.Info("established connection to db")
	return
}

func migrateUp() error {
	src, err := iofs.New(MigrationFS, "sql/migrations")
	if err != nil {
		return fmt.Errorf("failed to create iofs: %w", err)
	}
	m, err := migrate.NewWithSourceInstance(
		"iofs",
		src,
		fmt.Sprintf(
			"pgx5://%s:%s@%s/%s",
			os.Getenv("PGUSER"),
			os.Getenv("PGPASSWORD"),
			os.Getenv("PGHOST"),
			os.Getenv("PGDATABASE"),
		),
	)
	if err != nil {
		return fmt.Errorf("failed to set up migration: %w", err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to migrate db: %w", err)
	}
	return nil
}
