package http_operations

import "errors"

var (
	ErrOperationAlreadyExists = errors.New("operation (project + http method + route) already exists")
	ErrHTTPOperationNotFound  = errors.New("http operation not found")
)
