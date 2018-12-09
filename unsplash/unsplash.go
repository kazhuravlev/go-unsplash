package unsplash

import (
	"errors"
	"net/http"
)

var (
	ErrUnexpected   = errors.New("unexpected error")
	ErrBadRequest   = errors.New("bad request")
	ErrUnauthorized = errors.New("unauthorized")
)

const (
	apiURL = "https://api.unsplash.com"

	maxListItems = 30

	paginatorHeaderPerPage = "X-Per-Page"
	paginatorHeaderTotal   = "X-Total"

	rateLimitHeaderTotal     = "X-Ratelimit-Limit"
	rateLimitHeaderRemaining = "X-Ratelimit-Remaining"
)

func handleError(resp *http.Response) error {
	switch resp.StatusCode {
	case http.StatusBadRequest:
		return ErrBadRequest
	case http.StatusUnauthorized:
		return ErrUnauthorized
	default:
		return ErrUnexpected
	}
}
