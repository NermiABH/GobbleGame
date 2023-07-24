package vars

import "errors"

var (
	ErrMethodNotAllowed     = errors.New("method is not allowed")
	ErrInternalServer       = errors.New("internal server error")
	ErrMalformedContentType = errors.New("malformed Content-Type header")
	ErrUnsupportedMediaType = errors.New("Content-Type header must be application/json")
)
