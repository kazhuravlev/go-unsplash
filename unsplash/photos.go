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

func (o GetRandomPhotosOptions) validate() error {
	if o.Count > maxListItems {
		return ErrBadRequest
	}

	if o.Count < 0 {
		return ErrBadRequest
	}

	switch o.Orientation {
	case "", OrientationLandscape, OrientationPortrait, OrientationSquarish:
	default:
		return ErrBadRequest
	}

	return nil
}

func (o GetRandomPhotosOptions) query() url.Values {
	if o.Count == 0 {
		o.Count = 1
	}

	query := url.Values{}
	if len(o.Collections) != 0 {
		query.Set("collections", strings.Join(o.Collections, ","))
	}

	if o.Featured != "" {
		query.Set("featured", o.Featured)
	}

	if o.Username != "" {
		query.Set("username", o.Username)
	}

	if o.Query != "" {
		query.Set("query", o.Query)
	}

	if o.Orientation != "" {
		query.Set("orientation", string(o.Orientation))
	}

	query.Set("count", strconv.Itoa(o.Count))

	return query
}

func (c *Client) GetRandomPhotos(ctx context.Context, opts GetRandomPhotosOptions) ([]Photo, error) {
	if err := opts.validate(); err != nil {
		return nil, err
	}

	u := apiURL + "/photos/random?" + opts.query().Encode()

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

type OrderBy string

const (
	OrderByLatest  OrderBy = "latest"
	OrderByOldest  OrderBy = "oldest"
	OrderByPopular OrderBy = "popular"
)

type GetPhotosOptions struct {
	// Page is Page number to retrieve
	Page int
	// Per_page is Number of items per page
	PerPage int
	// Order_by is How to sort the photos
	OrderBy OrderBy
}

func (o GetPhotosOptions) validate() error {
	if o.Page < 0 {
		return ErrBadRequest
	}

	if o.PerPage < 0 {
		return ErrBadRequest
	}

	if o.PerPage > maxListItems {
		return ErrBadRequest
	}

	return nil
}

func (o GetPhotosOptions) query() url.Values {
	query := url.Values{}
	if o.Page == 0 {
		o.Page = 1
	}

	if o.PerPage == 0 {
		o.PerPage = 10
	}

	if o.OrderBy == "" {
		o.OrderBy = OrderByPopular
	}

	query.Set("page", strconv.Itoa(o.Page))
	query.Set("per_page", strconv.Itoa(o.PerPage))
	query.Set("order_by", string(o.OrderBy))

	return query
}

func (c *Client) GetPhotos(ctx context.Context, opts GetPhotosOptions) ([]Photo, error) {
	if err := opts.validate(); err != nil {
		return nil, err
	}

	u := apiURL + "/photos?" + opts.query().Encode()

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

func (c *Client) GetCuratedPhotos(ctx context.Context, opts GetPhotosOptions) ([]Photo, error) {
	if err := opts.validate(); err != nil {
		return nil, err
	}

	u := apiURL + "/photos/curated?" + opts.query().Encode()

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

func (c *Client) GetPhoto(ctx context.Context, id string) (*Photo, error) {
	if id == "" {
		return nil, ErrBadRequest
	}

	u := apiURL + "/photos/" + id

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

	var photo Photo
	if err := json.NewDecoder(resp.Body).Decode(&photo); err != nil {
		return nil, err
	}

	return &photo, nil
}
