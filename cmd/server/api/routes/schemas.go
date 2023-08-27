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

type HTTPOperation struct {
	ID             string    `json:"id"`
	ProjectID      string    `json:"project_id"`
	Method         string    `json:"method"`
	MIMEType       string    `json:"mime_type"`
	Route          string    `json:"route"`
	ResponseStatus uint16    `json:"response_status"`
	ResponseBody   string    `json:"response_body"`
	CreateAt       time.Time `json:"create_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func fromHTTPOperationDomain(operation entities.HTTPOperation) HTTPOperation {
	return HTTPOperation{
		ID:             operation.ID.String(),
		ProjectID:      operation.ProjectID.String(),
		Method:         string(operation.Method),
		MIMEType:       string(operation.MIMEType),
		Route:          operation.Route,
		ResponseStatus: uint16(operation.ResponseStatus),
		ResponseBody:   operation.ResponseBody,
		CreateAt:       operation.CreateAt,
		UpdatedAt:      operation.UpdatedAt,
	}
}
