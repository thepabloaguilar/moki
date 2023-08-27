package entities

import "errors"

var ErrInvalidHTTPMethod = errors.New("string doesn't match any http method")

type HTTPMethod string

const (
	HTTPMethodGet     HTTPMethod = "GET"
	HTTPMethodHead    HTTPMethod = "HEAD"
	HTTPMethodPost    HTTPMethod = "POST"
	HTTPMethodPut     HTTPMethod = "PUT"
	HTTPMethodPatch   HTTPMethod = "PATCH"
	HTTPMethodDelete  HTTPMethod = "DELETE"
	HTTPMethodConnect HTTPMethod = "CONNECT"
	HTTPMethodOptions HTTPMethod = "OPTIONS"
	HTTPMethodTrace   HTTPMethod = "TRACE"
)

var httpMethodByString = map[string]HTTPMethod{
	"GET":     HTTPMethodGet,
	"HEAD":    HTTPMethodHead,
	"POST":    HTTPMethodPost,
	"PUT":     HTTPMethodPut,
	"PATCH":   HTTPMethodPatch,
	"DELETE":  HTTPMethodDelete,
	"CONNECT": HTTPMethodConnect,
	"OPTIONS": HTTPMethodOptions,
	"TRACE":   HTTPMethodTrace,
}

func HTTPMethodFromString(str string) (HTTPMethod, error) {
	method, found := httpMethodByString[str]
	if !found {
		return "", ErrInvalidHTTPMethod
	}

	return method, nil
}
