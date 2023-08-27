package test_resources

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"

	"github.com/thepabloaguilar/moki/gateways/postgres"
)

const fiveMinutesExpiration = 300

type PostgresResource struct {
	connPool *pgxpool.Pool
}

func (pr *PostgresResource) NewDB(t *testing.T, ctx context.Context) *pgxpool.Pool {
	t.Helper()

	dbName := fmt.Sprintf(
		"moki_%s",
		strings.ReplaceAll(uuid.NewString(), "-", "_"),
	)
	createDBSQL := fmt.Sprintf("CREATE DATABASE %s TEMPLATE moki_template", dbName)

	_, err := pr.connPool.Exec(ctx, createDBSQL)
	require.NoError(t, err)

	dbAddress := strings.Replace(pr.connPool.Config().ConnString(), pr.connPool.Config().ConnConfig.Database, dbName, 1)
	newConnPool, err := pgxpool.New(ctx, dbAddress)
	require.NoError(t, err)

	t.Cleanup(func() {
		newConnPool.Close()

		_, err := pr.connPool.Exec(ctx, fmt.Sprintf("DROP DATABASE %s", dbName))
		require.NoError(t, err)
	})

	return newConnPool
}

func NewPostgresContainer(ctx context.Context, pool *dockertest.Pool, migrationsPath string) (PostgresResource, *dockertest.Resource, error) {
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "15-alpine",
		Env: []string{
			"POSTGRES_PASSWORD=postgres",
			"POSTGRES_DB=moki_test",
			"POSTGRES_USER=user",
			"POSTGRES_PASSWORD=password",
			"POSTGRES_HOST=postgres",
		},
	}, withAutoRemove)
	if err != nil {
		return PostgresResource{}, nil, err
	}

	if err = resource.Expire(fiveMinutesExpiration); err != nil {
		return PostgresResource{}, nil, err
	}

	dbAddress := fmt.Sprintf(
		"postgres://user:password@localhost:%s/moki_test?sslmode=disable",
		resource.GetPort("5432/tcp"),
	)

	err = pool.Retry(postgresHealthCheck(ctx, dbAddress))
	if err != nil {
		return PostgresResource{}, nil, err
	}

	pgConn, err := postgres.NewConnectionPool(ctx, dbAddress, 1, 4)
	if err != nil {
		return PostgresResource{}, nil, err
	}

	_, err = pgConn.Exec(ctx, "CREATE DATABASE moki_template")
	if err != nil {
		return PostgresResource{}, nil, err
	}

	dbTemplateAddress := strings.Replace(
		pgConn.Config().ConnString(),
		pgConn.Config().ConnConfig.Database,
		"moki_template",
		1,
	)
	if err := runMigrations(migrationsPath, dbTemplateAddress); err != nil {
		return PostgresResource{}, nil, err
	}

	return PostgresResource{
		connPool: pgConn,
	}, resource, nil
}

func withAutoRemove(cfg *docker.HostConfig) {
	cfg.AutoRemove = true
	cfg.RestartPolicy = docker.RestartPolicy{
		Name: "no",
	}
}

func postgresHealthCheck(ctx context.Context, dbAddress string) func() error {
	return func() error {
		connPool, err := postgres.NewConnectionPool(ctx, dbAddress, 1, 4)
		if err != nil {
			return err
		}
		defer connPool.Close()

		err = connPool.AcquireFunc(ctx, func(conn *pgxpool.Conn) error {
			return conn.Ping(ctx)
		})
		if err != nil {
			return err
		}

		return nil
	}
}

func runMigrations(migrationsPath string, dbAddress string) error {
	migration, err := migrate.New(fmt.Sprintf("file://%s", migrationsPath), dbAddress)
	if err != nil {
		return err
	}

	if err := migration.Up(); err != nil {
		return err
	}

	sourceErr, dbErr := migration.Close()
	if sourceErr != nil {
		return sourceErr
	}

	if dbErr != nil {
		return dbErr
	}

	return nil
}
