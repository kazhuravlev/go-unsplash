package unsplash

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type GetRandomPhotosOptions struct {
	// Public collection ID(‘s) to filter selection
	Collections []string
	// Limit selection to featured photos.
	Featured string
	// Limit selection to a single user.
	Username string
	// Limit selection to photos matching a search term.
	Query string
	// Filter search results by photo orientation. Valid values are landscape, portrait, and squarish.
	Orientation Orientation
	// The number of photos to return. (Default: 1; max: 30)
	Count int
}

func (c *Client) GetRandomPhotos(ctx context.Context, opts GetRandomPhotosOptions) ([]Photo, error) {
	if opts.Count > maxListItems {
		return nil, ErrBadRequest
	}

	if opts.Count < 0 {
		return nil, ErrBadRequest
	}

	query := url.Values{}
	query.Set("collections", strings.Join(opts.Collections, ","))
	query.Set("featured", opts.Featured)
	query.Set("username", opts.Username)
	query.Set("query", opts.Query)
	query.Set("orientation", string(opts.Orientation))
	query.Set("count", strconv.Itoa(opts.Count))

	u := apiURL + "/photos/random?" + query.Encode()

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, handleError(resp)
	}

	var photos []Photo
	if err := json.NewDecoder(resp.Body).Decode(&photos); err != nil {
		return nil, err
	}

	return photos, nil
}
