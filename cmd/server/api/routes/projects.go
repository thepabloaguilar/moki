package routes

import (
	"context"
	"errors"
	"net/http"

	"github.com/thepabloaguilar/moki/cmd/server/api"

	"github.com/go-chi/chi/v5"

	"github.com/thepabloaguilar/moki/core/projects"
	"github.com/thepabloaguilar/moki/extensions/chijson"
)

type ProjectsUseCases interface {
	CreateProject(ctx context.Context, input projects.CreateProjectInput) (projects.CreateProjectOutput, error)
}

func Projects(ucs ProjectsUseCases) func(router chi.Router) {
	return func(r chi.Router) {
		r.Post("/", chijson.Handler(createProject(ucs), chijson.WithSuccessStatus(http.StatusCreated)))
	}
}

type createProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func createProject(ucs ProjectsUseCases) chijson.HandlerFunc[createProjectRequest] {
	return func(req createProjectRequest, r *http.Request) (any, error) {
		output, err := ucs.CreateProject(r.Context(), projects.CreateProjectInput{
			ProjectName:        req.Name,
			ProjectDescription: req.Description,
		})
		if err != nil {
			if errors.Is(err, projects.ErrEmptyProjectName) {
				return nil, api.NewBadRequest(err)
			}

			return nil, err
		}

		return projectFromDomain(output.CreatedProject), err
	}
}
