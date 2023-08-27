package main

import (
	"errors"
	"flag"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	database = flag.String(
		"database",
		"postresql://moki:moki@locahost:5432/moki?sslmode=disable",
		"PostgreSQL Database connection string",
	)
	migrationsPath = flag.String(
		"migrations",
		"file://gateways/postgres/migrations",
		"Migrations path",
	)
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("error executing the migrations: %s", err)
	}
}

func run() error {
	flag.Parse()

	migration, err := migrate.New(*migrationsPath, *database)
	if err != nil {
		return err
	}

	err = migration.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
