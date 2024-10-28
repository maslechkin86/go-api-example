package port

import "errors"

var (
	ErrInvalidHTTPPort   = errors.New("invalid http port")
	ErrNoAuthHeader      = errors.New("no auth header")
	ErrInvalidAuthHeader = errors.New("invalid auth header")
	ErrGetBearerToken    = errors.New("unable to get bearer token")
)
