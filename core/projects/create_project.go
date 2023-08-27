package projects

import (
	"context"
	"strings"

	"github.com/thepabloaguilar/moki/core"
	"github.com/thepabloaguilar/moki/core/entities"
)

type CreateProjectUseCase struct {
	now         core.Now
	uuid        core.UUID
	persistence CreateProjectPersistenceGateway
}

type CreateProjectPersistenceGateway interface {
	CreateProject(ctx context.Context, project entities.Project) (entities.Project, error)
}

func NewCreateProject(now core.Now, uuid core.UUID, persistence CreateProjectPersistenceGateway) *CreateProjectUseCase {
	return &CreateProjectUseCase{
		now:         now,
		uuid:        uuid,
		persistence: persistence,
	}
}

type CreateProjectInput struct {
	ProjectName        string
	ProjectDescription string
}

type CreateProjectOutput struct {
	CreatedProject entities.Project
}

func (uc *CreateProjectUseCase) CreateProject(ctx context.Context, input CreateProjectInput) (CreateProjectOutput, error) {
	if err := uc.validateInput(input); err != nil {
		return CreateProjectOutput{}, err
	}

	createdProject, err := uc.persistence.CreateProject(ctx, entities.Project{
		ID:          uc.uuid(),
		Name:        input.ProjectName,
		Description: input.ProjectDescription,
		CreatedAt:   uc.now(),
		UpdatedAt:   uc.now(),
	})
	if err != nil {
		return CreateProjectOutput{}, err
	}

	return CreateProjectOutput{
		CreatedProject: createdProject,
	}, nil
}

func (uc *CreateProjectUseCase) validateInput(input CreateProjectInput) error {
	switch {
	case strings.Trim(input.ProjectName, " ") == "":
		return ErrEmptyProjectName
	default:
		return nil
	}
}
