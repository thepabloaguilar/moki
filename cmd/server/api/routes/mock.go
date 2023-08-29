package routes

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/thepabloaguilar/moki/core/http_operations"
	"github.com/thepabloaguilar/moki/core/mock"
)

type MockUseCases interface {
	ExecuteMock(ctx context.Context, input mock.ExecuteMockInput) (mock.ExecuteMockOutput, error)
}

func Mock(r chi.Router, ucs MockUseCases) {
	r.HandleFunc("/{projectID}/*", handleMock(ucs))
}

func handleMock(ucs MockUseCases) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projectID := chi.URLParam(r, "projectID")
		parsedProjectID, err := uuid.Parse(projectID)
		if err != nil {
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(`{"code": 501, "message": "invalid project id"}`)) //nolint:errcheck
		}

		output, err := ucs.ExecuteMock(r.Context(), mock.ExecuteMockInput{
			ProjectID:  parsedProjectID,
			HTTPMethod: r.Method,
			Route:      strings.TrimPrefix(r.URL.Path, fmt.Sprintf("/mock/%s", projectID)),
		})
		if err != nil {
			if errors.Is(err, http_operations.ErrHTTPOperationNotFound) {
				w.WriteHeader(http.StatusNotImplemented)
				w.Write([]byte(`{"code": 501, "message": "operation not found"}`)) //nolint:errcheck
				return
			}

			var noMatchingErr *mock.NotFoundMatchingOperationError
			if errors.As(err, &noMatchingErr) {
				w.WriteHeader(http.StatusNotImplemented)
				w.Write([]byte(fmt.Sprintf(`{"code": 501, "message": "%s"}`, noMatchingErr.Error()))) //nolint:errcheck
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", string(output.ResponseMimeType))
		w.WriteHeader(int(output.ResponseStatus))
		w.Write([]byte(output.ResponseBody)) //nolint:errcheck
	}
}
