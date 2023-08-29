package routes

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"

	"github.com/thepabloaguilar/moki/core/entities"
	"github.com/thepabloaguilar/moki/core/http_operations"

	"github.com/go-chi/chi/v5"

	"github.com/thepabloaguilar/moki/core/projects"
)

type ProjectsUseCases interface {
	CreateProject(ctx context.Context, input projects.CreateProjectInput) (projects.CreateProjectOutput, error)
	CreateHTTPOperation(ctx context.Context, input http_operations.CreateHTTPOperationInput) (http_operations.CreateHTTPOperationOutput, error)
}

func Projects(r chi.Router, ucs ProjectsUseCases) {
	r.Post("/projects", JSONHandler(createProject(ucs), WithSuccessStatus(http.StatusCreated)))
	r.Post("/projects/{projectID}/http-operations", JSONHandler(createHTTPOperation(ucs), WithSuccessStatus(http.StatusCreated)))
}

type createProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func createProject(ucs ProjectsUseCases) HandlerFunc[createProjectRequest] {
	return func(req createProjectRequest, r *http.Request) (any, error) {
		output, err := ucs.CreateProject(r.Context(), projects.CreateProjectInput{
			ProjectName:        req.Name,
			ProjectDescription: req.Description,
		})
		if err != nil {
			if errors.Is(err, projects.ErrEmptyProjectName) {
				return nil, NewBadRequest(err)
			}

			return nil, err
		}

		return projectFromDomain(output.CreatedProject), err
	}
}

type createHTTPOperationRequest struct {
	Method         string `json:"method"`
	MIMEType       string `json:"mime_type"`
	Route          string `json:"route"`
	ResponseStatus uint16 `json:"response_status"`
	ResponseBody   string `json:"response_body"`
}

func createHTTPOperation(ucs ProjectsUseCases) HandlerFunc[createHTTPOperationRequest] {
	return func(req createHTTPOperationRequest, r *http.Request) (any, error) {
		projectID := chi.URLParam(r, "projectID")
		parsedProjectID, err := uuid.Parse(projectID)
		if err != nil {
			return nil, NewBadRequest(errors.New("invalid project id"))
		}

		output, err := ucs.CreateHTTPOperation(r.Context(), http_operations.CreateHTTPOperationInput{
			ProjectID:      parsedProjectID,
			Method:         req.Method,
			MIMEType:       req.MIMEType,
			Route:          req.Route,
			ResponseStatus: req.ResponseStatus,
			ResponseBody:   req.ResponseBody,
		})
		if err != nil {
			errsIs := map[error]ErrorBuilder{
				entities.ErrInvalidHTTPMethod:             NewBadRequest,
				entities.ErrUnsupportedMIMEType:           NewBadRequest,
				entities.ErrInvalidHTTPStatusValue:        NewBadRequest,
				projects.ErrProjectNotFound:               NewPreconditionFailed,
				http_operations.ErrOperationAlreadyExists: NewPreconditionFailed,
			}
			for errIs, builder := range errsIs {
				if errors.Is(err, errIs) {
					return nil, builder(err)
				}
			}

			return nil, err
		}

		return fromHTTPOperationDomain(output.CreatedOperation), nil
	}
}
