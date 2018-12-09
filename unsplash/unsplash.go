package unsplash

import (
	"errors"
	"net/http"
	"strconv"
)

var (
	ErrUnexpected    = errors.New("unexpected error")
	ErrInvalidLimits = errors.New("invalid limits")
	ErrBadRequest    = errors.New("bad request")
	ErrUnauthorized  = errors.New("unauthorized")
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

type RateLimit struct {
	Limit     int
	Remaining int
}

func getLimits(resp *http.Response) (*RateLimit, error) {
	limit := resp.Header.Get(rateLimitHeaderTotal)
	if limit == "" {
		return nil, ErrInvalidLimits
	}

	remaining := resp.Header.Get(rateLimitHeaderRemaining)
	if remaining == "" {
		return nil, ErrInvalidLimits
	}

	limitInt, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		return nil, ErrInvalidLimits
	}

	remainingInt, err := strconv.ParseInt(remaining, 10, 64)
	if err != nil {
		return nil, ErrInvalidLimits
	}

	return &RateLimit{
		Limit:     int(limitInt),
		Remaining: int(remainingInt),
	}, nil
}
