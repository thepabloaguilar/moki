package entities

import (
	"time"

	"github.com/google/uuid"
)

type HTTPOperation struct {
	ID             uuid.UUID
	ProjectID      uuid.UUID
	Method         HTTPMethod
	MIMEType       MIMEType
	Route          string
	ResponseStatus HTTPStatus
	ResponseBody   string
	CreateAt       time.Time
	UpdatedAt      time.Time
}
