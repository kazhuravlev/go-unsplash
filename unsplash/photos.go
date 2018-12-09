package unsplash

import (
	"context"
	"encoding/json"
	"fmt"
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

	if o.Query != "" && len(o.Collections) != 0 {
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

	// always set count param. If this param do not present in query - response
	// will be an object, not list.
	query.Set("count", strconv.Itoa(o.Count))

	return query
}

func (c *Client) GetRandomPhotos(ctx context.Context, opts GetRandomPhotosOptions) ([]Photo, *RateLimit, error) {
	if err := opts.validate(); err != nil {
		return nil, nil, err
	}

	u := apiURL + "/photos/random?" + opts.query().Encode()

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	rl, err := getLimits(resp)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, rl, handleError(resp)
	}

	var photos []Photo
	if err := json.NewDecoder(resp.Body).Decode(&photos); err != nil {
		return nil, rl, err
	}

	return photos, rl, nil
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

func (c *Client) GetPhotos(ctx context.Context, opts GetPhotosOptions) ([]Photo, *RateLimit, error) {
	if err := opts.validate(); err != nil {
		return nil, nil, err
	}

	u := apiURL + "/photos?" + opts.query().Encode()

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	rl, err := getLimits(resp)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, rl, handleError(resp)
	}

	var photos []Photo
	if err := json.NewDecoder(resp.Body).Decode(&photos); err != nil {
		return nil, rl, err
	}

	return photos, rl, nil
}

func (c *Client) GetCuratedPhotos(ctx context.Context, opts GetPhotosOptions) ([]Photo, *RateLimit, error) {
	if err := opts.validate(); err != nil {
		return nil, nil, err
	}

	u := apiURL + "/photos/curated?" + opts.query().Encode()

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	rl, err := getLimits(resp)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, rl, handleError(resp)
	}

	var photos []Photo
	if err := json.NewDecoder(resp.Body).Decode(&photos); err != nil {
		return nil, rl, err
	}

	return photos, rl, nil
}

func (c *Client) GetPhoto(ctx context.Context, id string) (*Photo, *RateLimit, error) {
	if id == "" {
		return nil, nil, ErrBadRequest
	}

	u := apiURL + "/photos/" + id

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	rl, err := getLimits(resp)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, rl, handleError(resp)
	}

	var photo Photo
	if err := json.NewDecoder(resp.Body).Decode(&photo); err != nil {
		return nil, rl, err
	}

	return &photo, rl, nil
}

type Resolution string

const (
	ResolutionDays Resolution = "days"
)

type GetPhotoStatisticsOptions struct {
	// ID The public id of the photo
	ID string
	// Resolution The frequency of the stats
	Resolution Resolution
	// Quantity The amount of for each stat
	Quantity int
}

func (o GetPhotoStatisticsOptions) validate() error {
	if o.ID == "" {
		return ErrBadRequest
	}

	if o.Quantity < 0 {
		return ErrBadRequest
	}

	if o.Quantity > maxListItems {
		return ErrBadRequest
	}

	return nil
}

func (o GetPhotoStatisticsOptions) query() url.Values {
	query := url.Values{}

	if o.Resolution == "" {
		o.Resolution = ResolutionDays
	}

	if o.Quantity == 0 {
		o.Quantity = 1
	}

	query.Set("resolution", string(o.Resolution))
	query.Set("quantity", strconv.Itoa(o.Quantity))

	return query
}

func (c *Client) GetPhotoStatistics(ctx context.Context, opts GetPhotoStatisticsOptions) (*PhotoStatistics, *RateLimit, error) {
	if err := opts.validate(); err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("%s/photos/%s/statistics?%s", apiURL, opts.ID, opts.query().Encode())

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	rl, err := getLimits(resp)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, rl, handleError(resp)
	}

	var stat PhotoStatistics
	if err := json.NewDecoder(resp.Body).Decode(&stat); err != nil {
		return nil, rl, err
	}

	return &stat, rl, nil
}

func (c *Client) GetPhotoDownload(ctx context.Context, id string) (*PhotoDownload, *RateLimit, error) {
	if id == "" {
		return nil, nil, ErrBadRequest
	}

	u := fmt.Sprintf("%s/photos/%s/download", apiURL, id)

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	rl, err := getLimits(resp)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, rl, handleError(resp)
	}

	var download PhotoDownload
	if err := json.NewDecoder(resp.Body).Decode(&download); err != nil {
		return nil, rl, err
	}

	return &download, rl, nil
}

type Confidential int

const (
	ConfidentialYes Confidential = iota
	ConfidentialNo
)

type UpdateLocation struct {
	Latitude     float64
	Longitude    float64
	Name         string
	City         string
	Country      string
	Confidential Confidential
}

type UpdateExif struct {
	Make            string
	Models          string
	ExposureTime    string
	ApertureValue   string
	FocalLength     string
	ISOSpeedRatings string
}

type UpdatePhotoOptions struct {
	ID       string
	Location UpdateLocation
	Exif     UpdateExif
}

func (o UpdatePhotoOptions) validate() error {

	return nil
}

func (o UpdatePhotoOptions) query() url.Values {
	query := url.Values{}

	exif := o.Exif
	if exif.Make != "" {
		query.Set("exif[make]", exif.Make)
	}

	if exif.Models != "" {
		query.Set("exif[models]", exif.Models)
	}

	if exif.ExposureTime != "" {
		query.Set("exif[exposure_time]", exif.ExposureTime)
	}

	if exif.ApertureValue != "" {
		query.Set("exif[aperture_value]", exif.ApertureValue)
	}

	if exif.FocalLength != "" {
		query.Set("exif[focal_length]", exif.FocalLength)
	}

	if exif.ISOSpeedRatings != "" {
		query.Set("exif[iso_speed_ratings]", exif.ISOSpeedRatings)
	}

	location := o.Location
	if location.Latitude != 0 {
		query.Set("location[latitude]", strconv.FormatFloat(location.Latitude, 'f', 10, 64))
	}

	if location.Longitude != 0 {
		query.Set("location[longitude]", strconv.FormatFloat(location.Longitude, 'f', 10, 64))
	}

	if location.Name != "" {
		query.Set("location[name]", location.Name)
	}

	if location.City != "" {
		query.Set("location[city]", location.City)
	}

	if location.Country != "" {
		query.Set("location[country]", location.Country)
	}

	switch location.Confidential {
	case ConfidentialYes:
		query.Set("location[confidential]", "true")
	case ConfidentialNo:
		query.Set("location[confidential]", "false")
	}

	return query
}

func (c *Client) UpdatePhoto(ctx context.Context, opts UpdatePhotoOptions) (*Photo, *RateLimit, error) {
	if err := opts.validate(); err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("%s/photos/%s?%s", apiURL, opts.ID, opts.query().Encode())

	req, err := http.NewRequest(http.MethodPut, u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	rl, err := getLimits(resp)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, rl, handleError(resp)
	}

	var photo Photo
	if err := json.NewDecoder(resp.Body).Decode(&photo); err != nil {
		return nil, rl, err
	}

	return &photo, rl, nil
}

func (c *Client) LikePhoto(ctx context.Context, id string) (*Photo, *RateLimit, error) {
	if id == "" {
		return nil, nil, ErrBadRequest
	}

	u := fmt.Sprintf("%s/photos/%s/like", apiURL, id)

	req, err := http.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	rl, err := getLimits(resp)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, rl, handleError(resp)
	}

	var photo Photo
	if err := json.NewDecoder(resp.Body).Decode(&photo); err != nil {
		return nil, rl, err
	}

	return &photo, rl, nil
}

func (c *Client) UnlikePhoto(ctx context.Context, id string) (*Photo, *RateLimit, error) {
	if id == "" {
		return nil, nil, ErrBadRequest
	}

	u := fmt.Sprintf("%s/photos/%s/like", apiURL, id)

	req, err := http.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	rl, err := getLimits(resp)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, rl, handleError(resp)
	}

	var photo Photo
	if err := json.NewDecoder(resp.Body).Decode(&photo); err != nil {
		return nil, rl, err
	}

	return &photo, rl, nil
}

type SearchPhotosOptions struct {
	// Query Search terms.
	Query string
	// Page Page number to retrieve. (Optional; default: 1)
	Page int
	// Per_page Number of items per page. (Optional; default: 10)
	PerPage int
	// Collections Collection ID(‘s) to narrow search. If multiple, comma-separated.
	Collections []string
	// Orientation Filter search results by photo orientation. Valid values are landscape, portrait, and squarish.
	Orientation Orientation
}

func (o SearchPhotosOptions) validate() error {
	if o.Page < 0 {
		return ErrBadRequest
	}

	if o.PerPage < 0 {
		return ErrBadRequest
	}

	if o.PerPage > maxListItems {
		return ErrBadRequest
	}

	switch o.Orientation {
	case "", OrientationLandscape, OrientationPortrait, OrientationSquarish:
	default:
		return ErrBadRequest
	}

	if o.Query != "" && len(o.Collections) != 0 {
		return ErrBadRequest
	}

	return nil
}

func (o SearchPhotosOptions) query() url.Values {
	query := url.Values{}
	if o.Page == 0 {
		o.Page = 1
	}

	if o.PerPage == 0 {
		o.PerPage = 10
	}

	if len(o.Collections) != 0 {
		query.Set("collections", strings.Join(o.Collections, ","))
	}

	if o.Query != "" {
		query.Set("query", o.Query)
	}

	if o.Orientation != "" {
		query.Set("orientation", string(o.Orientation))
	}

	query.Set("page", strconv.Itoa(o.Page))
	query.Set("per_page", strconv.Itoa(o.PerPage))

	return query
}

func (c *Client) SearchPhotos(ctx context.Context, opts SearchPhotosOptions) (*SearchResult, *RateLimit, error) {
	if err := opts.validate(); err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("%s/search/photos?%s", apiURL, opts.query().Encode())

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	rl, err := getLimits(resp)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, rl, handleError(resp)
	}

	var searchRes SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&searchRes); err != nil {
		return nil, rl, err
	}

	return &searchRes, rl, nil
}

type SearchCollectionsOptions struct {
	Query   string
	PerPage int
	Page    int
}

func (o SearchCollectionsOptions) validate() error {
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

func (o SearchCollectionsOptions) query() url.Values {
	query := url.Values{}
	if o.Page == 0 {
		o.Page = 1
	}

	if o.PerPage == 0 {
		o.PerPage = 10
	}

	if o.Query != "" {
		query.Set("query", o.Query)
	}

	query.Set("page", strconv.Itoa(o.Page))
	query.Set("per_page", strconv.Itoa(o.PerPage))

	return query
}

func (c *Client) SearchCollections(ctx context.Context, opts SearchCollectionsOptions) (*CollectionSearchResult, *RateLimit, error) {
	if err := opts.validate(); err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("%s/search/collections?%s", apiURL, opts.query().Encode())

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	rl, err := getLimits(resp)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, rl, handleError(resp)
	}

	var searchRes CollectionSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&searchRes); err != nil {
		return nil, rl, err
	}

	return &searchRes, rl, nil
}
