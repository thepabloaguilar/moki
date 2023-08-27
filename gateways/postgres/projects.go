package postgres

import (
	"context"
	"fmt"

	"github.com/thepabloaguilar/moki/core/entities"
)

const projectFields = "id, name, description, created_at, updated_at"

type Projects struct {
	db db
}

func NewProjects(db db) *Projects {
	return &Projects{db: db}
}

func scanProject(scanner scan) (entities.Project, error) {
	project := entities.Project{}
	err := scanner(
		&project.ID,
		&project.Name,
		&project.Description,
		&project.CreatedAt,
		&project.UpdatedAt,
	)
	if err != nil {
		return entities.Project{}, err
	}

	return project, err
}

func (p *Projects) CreateProject(ctx context.Context, project entities.Project) (entities.Project, error) {
	query := fmt.Sprintf(
		`INSERT INTO projects(%s) VALUES ($1, $2, $3, $4, $5) RETURNING %s`,
		projectFields, projectFields,
	)

	createdProject, err := scanProject(p.db.QueryRow(
		ctx, query, project.ID, project.Name, project.Description,
		project.CreatedAt, project.UpdatedAt,
	).Scan)
	if err != nil {
		return entities.Project{}, err
	}

	return createdProject, nil
}
