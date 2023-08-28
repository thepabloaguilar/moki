package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/thepabloaguilar/moki/core/mock"

	"github.com/thepabloaguilar/moki/core/http_operations"

	"github.com/rs/zerolog"

	"github.com/thepabloaguilar/moki/cmd/server/api/routes"

	"github.com/google/uuid"

	"github.com/thepabloaguilar/moki/core/projects"
	"github.com/thepabloaguilar/moki/gateways/postgres"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "go.uber.org/automaxprocs"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// TODO: Change the argument to `os.Stdout` when releasing the first version
	logger := zerolog.New(zerolog.NewConsoleWriter()).With().Timestamp().Logger()

	if err := run(logger); err != nil {
		logger.Fatal().Err(err).Msg("error running application")
	}
}

func run(logger zerolog.Logger) error {
	ctx := context.Background()
	r := chi.NewRouter()

	// Postgres
	pgPool, err := postgres.NewConnectionPool(ctx, "postgresql://moki:moki@localhost:5432/moki?sslmode=disable", 1, 2)
	if err != nil {
		return err
	}
	defer pgPool.Close()

	projectsRepo := postgres.NewProjects(pgPool)
	httpOperationsRepo := postgres.NewHTTPOperations(pgPool)

	now := func() time.Time {
		return time.Now().UTC()
	}

	// Use Cases
	createProject := projects.NewCreateProject(now, uuid.New, projectsRepo)

	createHTTPOperation := http_operations.NewCreateHTTPOperation(now, uuid.New, &struct {
		*postgres.Projects
		*postgres.HTTPOperations
	}{
		Projects:       projectsRepo,
		HTTPOperations: httpOperationsRepo,
	})

	executeMock := mock.NewExecuteMock(&struct {
		*postgres.Projects
		*postgres.HTTPOperations
	}{
		Projects:       projectsRepo,
		HTTPOperations: httpOperationsRepo,
	})

	// Use Cases Collection
	projectsCollection := &struct {
		*projects.CreateProjectUseCase
		*http_operations.CreateHTTPOperationUseCase
	}{
		CreateProjectUseCase:       createProject,
		CreateHTTPOperationUseCase: createHTTPOperation,
	}

	r.Use(middleware.Heartbeat("/ping"))

	r.Group(func(r chi.Router) {
		r.Use(
			middleware.Logger,
			middleware.Recoverer,
			middleware.AllowContentType("application/json"),
		)

		r.Route("/api", func(r chi.Router) {
			routes.Projects(r, projectsCollection)
		})
	})

	r.Group(func(r chi.Router) {
		r.Use(
			middleware.Logger,
			middleware.Recoverer,
		)
		routes.Mock(r, executeMock)
	})

	server := http.Server{
		Addr:              ":8000",
		Handler:           r,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	ctx, cancel := context.WithCancel(ctx)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-signals
		log.Printf("received signal %s", sig)

		timeoutCtx, cancelTimeoutCtx := context.WithTimeout(ctx, 10*time.Second)

		go func() {
			<-timeoutCtx.Done()

			if errors.Is(timeoutCtx.Err(), context.DeadlineExceeded) {
				logger.Fatal().Msg("forcing application shutdown")
			}
		}()

		pgPool.Close()

		if err := server.Shutdown(timeoutCtx); err != nil {
			logger.Error().Err(err).Msg("error shutingdown server")
		}

		cancel()
		cancelTimeoutCtx()
	}()

	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
