package mock

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/thepabloaguilar/moki/core/entities"
)

type NotFoundMatchingOperationError struct {
	projectID  uuid.UUID
	httpMethod entities.HTTPMethod
	route      string
}

func NewNotFoundMatchingOperationError(
	projectID uuid.UUID,
	httpMethod entities.HTTPMethod,
	route string,
) *NotFoundMatchingOperationError {
	return &NotFoundMatchingOperationError{
		projectID:  projectID,
		httpMethod: httpMethod,
		route:      route,
	}
}

func (e *NotFoundMatchingOperationError) Error() string {
	return fmt.Sprintf(
		"moki was unable to find a operation matching for project '%s': %s %s",
		e.projectID, e.httpMethod, e.route,
	)
}
