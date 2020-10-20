package services

import (
	"context"
	"os"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"

	models "github.com/nvonpentz/go-graphql-react/internal/models/db"
	// Needed for migration
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Config is configuration of the
type PostgresConfig struct {
	URL     string
	Migrate bool
}

type Postgres struct {
	*sqlx.DB
}

// New creates a new Postgres
func NewPostgres(config PostgresConfig) (*Postgres, error) {
	databaseURL := config.URL
	performMigration := config.Migrate

	db, err := sqlx.Open("postgres", databaseURL)
	if err != nil {
		return &Postgres{}, err
	}

	pg := &Postgres{db}

	if performMigration {
		err := pg.RunMigration()
		if err != nil {
			return &Postgres{}, err
		}
	}

	return pg, nil
}

func NewPostgresWithDefaults() (*Postgres, error) {
	postgres, err := NewPostgres(
		PostgresConfig{
			URL:     os.Getenv("POSTGRES_URL"),
			Migrate: true,
		},
	)
	if err != nil {
		return &Postgres{}, err
	}
	return postgres, nil
}

// NewMigration runs another migration
func (pg *Postgres) NewMigration() (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(pg.DB.DB, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	dbMigrationsURL := os.Getenv("POSTGRES_MIGRATIONS_FILE_URL")
	migration, err := migrate.NewWithDatabaseInstance(
		dbMigrationsURL,
		"postgres",
		driver,
	)
	if err != nil {
		return nil, err
	}

	return migration, err
}

// RunMigration runs migrations
func (pg *Postgres) RunMigration() error {
	m, err := pg.NewMigration()
	if err != nil {
		return err
	}

	migrateVersion := uint(1)
	err = m.Migrate(migrateVersion)
	if err != migrate.ErrNoChange && err != nil {
		return err
	}

	return nil
}

func (pg *Postgres) ClearTables() error {
	_, err := models.Users().DeleteAll(context.Background(), pg)
	return err
}
