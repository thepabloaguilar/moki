package routes

import (
	"time"

	"github.com/thepabloaguilar/moki/core/entities"
)

type Project struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func projectFromDomain(project entities.Project) Project {
	return Project{
		ID:          project.ID.String(),
		Name:        project.Name,
		Description: project.Description,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	}
}
