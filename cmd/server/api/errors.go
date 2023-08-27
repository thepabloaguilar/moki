package api

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	code int
	err  error
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("http error (%d): %s", e.code, e.err)
}

func NewBadRequest(err error) error {
	return &HttpError{
		code: http.StatusBadRequest,
		err:  err,
	}
}
