package mock_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"

	"github.com/thepabloaguilar/moki/extensions/test_resources"
)

const migrationPath = "../../gateways/postgres/migrations"

var postgresResource test_resources.PostgresResource

func TestMain(m *testing.M) {
	ctx, cancel := context.WithCancel(context.Background())
	dockerPoll, err := dockertest.NewPool("")
	if err != nil {
		cancel()
		log.Fatalf("error creating docker pool: %s", err)
	}

	pgResource, postgresDocker, err := test_resources.NewPostgresContainer(ctx, dockerPoll, migrationPath)
	if err != nil {
		cancel()
		log.Fatalf("error creating postgres resource: %s", err)
	}

	postgresResource = pgResource

	code := m.Run()

	if err := postgresDocker.Close(); err != nil {
		cancel()
		log.Fatalf("error purging postgres: %s", err)
	}

	cancel()

	os.Exit(code)
}
