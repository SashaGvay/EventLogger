package root

import (
	"errors"
	"fmt"
	"log"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx/v5/stdlib"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (r *Root) RunMigrations() error {
	db, err := postgres.WithInstance(stdlib.OpenDB(*r.Infrastructure.DbInstance.Config()), &postgres.Config{})
	if err != nil {
		return fmt.Errorf("postgres.WithInstance: %w", err)
	}

	migrator, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", db)
	if err != nil {
		return fmt.Errorf("migrate.NewWithDatabaseInstance: %w", err)
	}

	if err := migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migrator.Up: %w", err)
	}

	log.Println("migrations done")
	return nil
}
