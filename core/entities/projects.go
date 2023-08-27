package entities

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID          uuid.UUID
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
