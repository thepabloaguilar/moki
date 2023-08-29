package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ErrorBuilder func(error) error

type HttpError struct {
	code int
	err  error
}

func (e *HttpError) GetCode() int {
	return e.code
}

func (e *HttpError) GetError() error {
	return e.err
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("http error (%d): %s", e.code, e.err)
}

func (e *HttpError) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"code":    e.GetCode(),
		"message": e.err.Error(),
	})
}

func NewBadRequest(err error) error {
	return &HttpError{
		code: http.StatusBadRequest,
		err:  err,
	}
}

func NewPreconditionFailed(err error) error {
	return &HttpError{
		code: http.StatusPreconditionFailed,
		err:  err,
	}
}

func NewInternalServerError() *HttpError {
	return &HttpError{
		code: http.StatusInternalServerError,
		err:  errors.New("internal server error"),
	}
}
