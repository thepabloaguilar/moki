package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/thepabloaguilar/moki/core/http_operations"

	"github.com/thepabloaguilar/moki/core/entities"
)

const httpOperationFields = "id, project_id, http_method, mime_type, route, response_status, response_body, created_at, updated_at"
const uniqueOperationConstraintName = "http_operations_project_id_http_method_route_key"

type HTTPOperations struct {
	db db
}

func NewHTTPOperations(db db) *HTTPOperations {
	return &HTTPOperations{db: db}
}

func scanHTTPOperation(scanner scan) (entities.HTTPOperation, error) {
	operation := entities.HTTPOperation{}
	err := scanner(
		&operation.ID,
		&operation.ProjectID,
		&operation.Method,
		&operation.MIMEType,
		&operation.Route,
		&operation.ResponseStatus,
		&operation.ResponseBody,
		&operation.CreateAt,
		&operation.UpdatedAt,
	)
	if err != nil {
		return entities.HTTPOperation{}, err
	}

	return operation, nil
}

func (o *HTTPOperations) CreateHTTPOperation(ctx context.Context, operation entities.HTTPOperation) (entities.HTTPOperation, error) {
	query := fmt.Sprintf(
		`INSERT INTO http_operations(%s) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING %s`,
		httpOperationFields, httpOperationFields,
	)

	createOperation, err := scanHTTPOperation(o.db.QueryRow(
		ctx, query, operation.ID, operation.ProjectID, operation.Method, operation.MIMEType, operation.Route,
		operation.ResponseStatus, operation.ResponseBody, operation.CreateAt, operation.UpdatedAt,
	).Scan)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.ConstraintName == uniqueOperationConstraintName {
				return entities.HTTPOperation{}, http_operations.ErrOperationAlreadyExists
			}
		}

		return entities.HTTPOperation{}, err
	}

	return createOperation, nil
}
