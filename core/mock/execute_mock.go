package mock

import (
	"context"
	"errors"

	"github.com/thepabloaguilar/moki/core/http_operations"

	"github.com/google/uuid"

	"github.com/thepabloaguilar/moki/core/entities"
)

//go:generate moq -stub -pkg mocks -out mocks/execute_mock_persistence_gateway.go . ExecuteMockPersistenceGateway

type ExecuteMockUseCase struct {
	persistence ExecuteMockPersistenceGateway
}

type ExecuteMockPersistenceGateway interface {
	GetProjectByID(ctx context.Context, projectID uuid.UUID) (entities.Project, error)
	GetHTTPOperationByProjectIDAndHTTPMethodAndRoute(
		ctx context.Context,
		projectID uuid.UUID,
		httpMethod entities.HTTPMethod,
		route string,
	) (entities.HTTPOperation, error)
}

func NewExecuteMock(persistence ExecuteMockPersistenceGateway) *ExecuteMockUseCase {
	return &ExecuteMockUseCase{persistence: persistence}
}

type ExecuteMockInput struct {
	ProjectID  uuid.UUID
	HTTPMethod string
	Route      string
}

type ExecuteMockOutput struct {
	ResponseBody     string
	ResponseMimeType entities.MIMEType
	ResponseStatus   entities.HTTPStatus
}

func (uc *ExecuteMockUseCase) ExecuteMock(ctx context.Context, input ExecuteMockInput) (ExecuteMockOutput, error) {
	project, err := uc.persistence.GetProjectByID(ctx, input.ProjectID)
	if err != nil {
		return ExecuteMockOutput{}, err
	}

	httpMethod, err := entities.HTTPMethodFromString(input.HTTPMethod)
	if err != nil {
		return ExecuteMockOutput{}, err
	}

	operation, err := uc.persistence.GetHTTPOperationByProjectIDAndHTTPMethodAndRoute(
		ctx, project.ID, httpMethod, input.Route,
	)
	if err != nil {
		if errors.Is(err, http_operations.ErrHTTPOperationNotFound) {
			return ExecuteMockOutput{}, NewNotFoundMatchingOperationError(
				input.ProjectID, httpMethod, input.Route,
			)
		}

		return ExecuteMockOutput{}, err
	}

	return ExecuteMockOutput{
		ResponseBody:     operation.ResponseBody,
		ResponseMimeType: operation.MIMEType,
		ResponseStatus:   operation.ResponseStatus,
	}, nil
}
