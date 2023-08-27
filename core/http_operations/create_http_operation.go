package http_operations

import (
	"context"

	"github.com/google/uuid"

	"github.com/thepabloaguilar/moki/core"
	"github.com/thepabloaguilar/moki/core/entities"
)

//go:generate moq -stub -pkg mocks -out mocks/create_http_operation_persistence_gateway.go . CreateHTTPOperationPersistenceGateway

type CreateHTTPOperationUseCase struct {
	now         core.Now
	uuid        core.UUID
	persistence CreateHTTPOperationPersistenceGateway
}

type CreateHTTPOperationPersistenceGateway interface {
	GetProjectByID(ctx context.Context, projectID uuid.UUID) (entities.Project, error)
	CreateHTTPOperation(ctx context.Context, operation entities.HTTPOperation) (entities.HTTPOperation, error)
}

func NewCreateHTTPOperation(
	now core.Now,
	uuid core.UUID,
	persistence CreateHTTPOperationPersistenceGateway,
) *CreateHTTPOperationUseCase {
	return &CreateHTTPOperationUseCase{
		now:         now,
		uuid:        uuid,
		persistence: persistence,
	}
}

type CreateHTTPOperationInput struct {
	ProjectID      uuid.UUID
	Method         string
	MIMEType       string
	Route          string
	ResponseStatus uint16
	ResponseBody   string
}

type CreateHTTPOperationOutput struct {
	CreatedOperation entities.HTTPOperation
}

func (uc *CreateHTTPOperationUseCase) CreateHTTPOperation(ctx context.Context, input CreateHTTPOperationInput) (CreateHTTPOperationOutput, error) {
	createOperation, err := uc.createOperation(ctx, input)
	if err != nil {
		return CreateHTTPOperationOutput{}, err
	}

	return CreateHTTPOperationOutput{
		CreatedOperation: createOperation,
	}, nil
}

func (uc *CreateHTTPOperationUseCase) createOperation(ctx context.Context, input CreateHTTPOperationInput) (entities.HTTPOperation, error) {
	project, err := uc.persistence.GetProjectByID(ctx, input.ProjectID)
	if err != nil {
		return entities.HTTPOperation{}, err
	}

	method, err := entities.HTTPMethodFromString(input.Method)
	if err != nil {
		return entities.HTTPOperation{}, err
	}

	mimeType, err := entities.MIMETypeFromString(input.MIMEType)
	if err != nil {
		return entities.HTTPOperation{}, err
	}

	responseStatus, err := entities.HTTPStatusFromInt(input.ResponseStatus)
	if err != nil {
		return entities.HTTPOperation{}, err
	}

	createdOperation, err := uc.persistence.CreateHTTPOperation(ctx, entities.HTTPOperation{
		ID:             uc.uuid(),
		ProjectID:      project.ID,
		Method:         method,
		MIMEType:       mimeType,
		Route:          input.Route,
		ResponseStatus: responseStatus,
		ResponseBody:   input.ResponseBody,
		CreateAt:       uc.now(),
		UpdatedAt:      uc.now(),
	})
	if err != nil {
		return entities.HTTPOperation{}, err
	}

	return createdOperation, nil
}
