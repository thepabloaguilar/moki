package chijson

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type HandlerFunc[T any] func(body T, r *http.Request) (any, error)

type HandlerOption func(cfg *HandlerConfig)

type HandlerConfig struct {
	SuccessStatusCode int
}

func Handler[T any](handler HandlerFunc[T], opts ...HandlerOption) http.HandlerFunc {
	cfg := HandlerConfig{
		SuccessStatusCode: http.StatusOK,
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; chartset=utf-8")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"error reading body"}`)) //nolint:errcheck
			return
		}

		var unmarshalledBody T
		if r.ContentLength != 0 {
			if err = json.Unmarshal(body, &unmarshalledBody); err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte(`{"message":"error unmarshalling json"}`)) //nolint:errcheck
				return
			}
		}

		result, err := handler(unmarshalledBody, r)
		if err != nil {
			log.Printf("error on request: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		jsonResponse, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(`{"message":"error marshaling json"}`)) //nolint:errcheck
			return
		}

		w.WriteHeader(cfg.SuccessStatusCode)
		w.Write(jsonResponse) //nolint:errcheck
	}
}