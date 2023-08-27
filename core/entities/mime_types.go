package entities

import "errors"

var ErrUnsupportedMIMEType = errors.New("string doesn't match any supported mime type")

type MIMEType string

const (
	MIMETypeJSON MIMEType = "application/json"
	MIMETypeXML  MIMEType = "application/xml"
	MIMETypeText MIMEType = "text/plain"
)

var mimeTypeByString = map[string]MIMEType{
	"application/json": MIMETypeJSON,
	"application/xml":  MIMETypeXML,
	"text/plain":       MIMETypeText,
}

func MIMETypeFromString(str string) (MIMEType, error) {
	mimeType, found := mimeTypeByString[str]
	if !found {
		return "", ErrUnsupportedMIMEType
	}

	return mimeType, nil
}
