package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/thepabloaguilar/moki/cmd/server/api/routes"
)

type UseCasesCollection struct {
	MockUseCases    routes.MockUseCases
	ProjectUseCases routes.ProjectsUseCases
}

type Config struct {
	Address      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	UseCases     UseCasesCollection
}

func New(cfg Config) http.Server {
	router := chi.NewRouter()

	// TODO: Add a better health check later to see if the database is responding
	router.Use(middleware.Heartbeat("/health/check"))

	registerAPIRoutes(router, cfg.UseCases)
	registerMockRoutes(router, cfg.UseCases)

	return http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}
}

func registerAPIRoutes(r chi.Router, ucs UseCasesCollection) {
	r.Group(func(r chi.Router) {
		r.Use(
			middleware.Logger,
			middleware.Recoverer,
			middleware.AllowContentType("application/json"),
			middleware.CleanPath,
			middleware.GetHead,
			middleware.RedirectSlashes,
		)

		r.Route("/api", func(r chi.Router) {
			routes.Projects(r, ucs.ProjectUseCases)
		})
	})
}

func registerMockRoutes(r chi.Router, ucs UseCasesCollection) {
	r.Group(func(r chi.Router) {
		r.Use(
			middleware.Logger,
			middleware.Recoverer,
			middleware.AllowContentType("application/json", "application/xml", "text/plain"),
		)

		r.Route("/mock", func(r chi.Router) {
			routes.Mock(r, ucs.MockUseCases)
		})
	})
}
