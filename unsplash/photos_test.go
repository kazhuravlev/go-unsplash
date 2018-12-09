package unsplash_test

import (
	"context"
	"fmt"
	"github.com/kazhuravlev/go-unsplash/unsplash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	minLimits = 50
)

func TestClient_GetRandomPhotos(t *testing.T) {
	c, err := unsplash.New(unsplash.WithHttpClient(httpClient))
	require.Nil(t, err)

	n := 30

	photos, rl, err := c.GetRandomPhotos(context.Background(), unsplash.GetRandomPhotosOptions{Orientation: unsplash.OrientationLandscape, Count: n})
	require.Nil(t, err)
	assert.NotNil(t, rl)
	assert.True(t, rl.Limit >= minLimits)
	assert.NotEqual(t, rl.Remaining, 0)

	assert.NotNil(t, photos)
	assert.Len(t, photos, n)

	fmt.Println(photos[0].Urls.Full)
}

func TestClient_GetPhotos(t *testing.T) {
	c, err := unsplash.New(unsplash.WithHttpClient(httpClient))
	require.Nil(t, err)

	n := 30
	photos, rl, err := c.GetPhotos(context.Background(), unsplash.GetPhotosOptions{PerPage: n})
	require.Nil(t, err)
	assert.NotNil(t, rl)
	assert.True(t, rl.Limit >= minLimits)
	assert.NotEqual(t, rl.Remaining, 0)

	assert.NotNil(t, photos)
	assert.Len(t, photos, n)

	fmt.Println(photos[0].Urls.Full)
}

func TestClient_GetCuratedPhotos(t *testing.T) {
	c, err := unsplash.New(unsplash.WithHttpClient(httpClient))
	require.Nil(t, err)

	n := 30
	photos, rl, err := c.GetCuratedPhotos(context.Background(), unsplash.GetPhotosOptions{PerPage: n})
	require.Nil(t, err)
	assert.NotNil(t, rl)
	assert.True(t, rl.Limit >= minLimits)
	assert.NotEqual(t, rl.Remaining, 0)

	assert.NotNil(t, photos)
	assert.Len(t, photos, n)

	fmt.Println(photos[0].Urls.Full)
}

func TestClient_GetPhoto(t *testing.T) {
	c, err := unsplash.New(unsplash.WithHttpClient(httpClient))
	require.Nil(t, err)

	id := "pnNR3P5m15s"
	photo, rl, err := c.GetPhoto(context.Background(), id)
	require.Nil(t, err)
	assert.NotNil(t, rl)
	assert.True(t, rl.Limit >= minLimits)
	assert.NotEqual(t, rl.Remaining, 0)

	assert.NotNil(t, photo)
	assert.Equal(t, photo.ID, id)
	assert.Equal(t, photo.Location.Name, "example")
}

func TestClient_GetPhotoStatistics(t *testing.T) {
	c, err := unsplash.New(unsplash.WithHttpClient(httpClient))
	require.Nil(t, err)

	id := "Mg0W1N_yDv0"
	photo, rl, err := c.GetPhotoStatistics(context.Background(), unsplash.GetPhotoStatisticsOptions{ID: id})
	require.Nil(t, err)
	assert.NotNil(t, rl)
	assert.True(t, rl.Limit >= minLimits)
	assert.NotEqual(t, rl.Remaining, 0)

	assert.NotNil(t, photo)
	assert.Equal(t, photo.ID, id)
}

func TestClient_GetPhotoDownload(t *testing.T) {
	c, err := unsplash.New(unsplash.WithHttpClient(httpClient))
	require.Nil(t, err)

	id := "Mg0W1N_yDv0"
	photo, rl, err := c.GetPhotoDownload(context.Background(), id)
	require.Nil(t, err)
	assert.NotNil(t, rl)
	assert.True(t, rl.Limit >= minLimits)
	assert.NotEqual(t, rl.Remaining, 0)

	assert.NotNil(t, photo)
	assert.Equal(t, photo.URL, "https://images.unsplash.com/photo-1536167038724-17be8c5e6876?ixlib=rb-1.2.1&q=85&fm=jpg&crop=entropy&cs=srgb")
}

func TestClient_UpdatePhoto(t *testing.T) {
	c, err := unsplash.New(unsplash.WithHttpClient(httpClient))
	require.Nil(t, err)

	id := "pnNR3P5m15s"
	photo, rl, err := c.UpdatePhoto(context.Background(), unsplash.UpdatePhotoOptions{ID: id, Location: unsplash.UpdateLocation{Name: "example", Longitude: 0.34}})
	require.Nil(t, err)
	assert.NotNil(t, rl)
	assert.True(t, rl.Limit >= minLimits)
	assert.NotEqual(t, rl.Remaining, 0)

	assert.NotNil(t, photo)
	assert.Equal(t, photo.Location.Name, "example")
}

func TestClient_LikePhoto(t *testing.T) {
	c, err := unsplash.New(unsplash.WithHttpClient(httpClient))
	require.Nil(t, err)

	id := "pnNR3P5m15s"
	photo, rl, err := c.LikePhoto(context.Background(), id)
	require.Nil(t, err)
	assert.NotNil(t, rl)
	assert.True(t, rl.Limit >= minLimits)
	assert.NotEqual(t, rl.Remaining, 0)

	assert.NotNil(t, photo)
	assert.Equal(t, photo.Location.Name, "example")
}

func TestClient_UnlikePhoto(t *testing.T) {
	c, err := unsplash.New(unsplash.WithHttpClient(httpClient))
	require.Nil(t, err)

	id := "pnNR3P5m15s"
	photo, rl, err := c.UnlikePhoto(context.Background(), id)
	require.Nil(t, err)
	assert.NotNil(t, rl)
	assert.True(t, rl.Limit >= minLimits)
	assert.NotEqual(t, rl.Remaining, 0)

	assert.NotNil(t, photo)
	assert.Equal(t, photo.Location.Name, "example")
}

func TestClient_SearchPhotos(t *testing.T) {
	c, err := unsplash.New(unsplash.WithHttpClient(httpClient))
	require.Nil(t, err)

	res, rl, err := c.SearchPhotos(context.Background(), unsplash.SearchPhotosOptions{Query: "car"})
	require.Nil(t, err)
	assert.NotNil(t, rl)
	assert.True(t, rl.Limit >= minLimits)
	assert.NotEqual(t, rl.Remaining, 0)

	assert.NotNil(t, res)
	assert.NotEqual(t, res.Total, 0)
	assert.NotEqual(t, res.TotalPages, 0)
	assert.NotEmpty(t, res.Results)
}

func TestClient_SearchCollections(t *testing.T) {
	c, err := unsplash.New(unsplash.WithHttpClient(httpClient))
	require.Nil(t, err)

	res, rl, err := c.SearchCollections(context.Background(), unsplash.SearchCollectionsOptions{Query: "car"})
	require.Nil(t, err)
	assert.NotNil(t, rl)
	assert.True(t, rl.Limit >= minLimits)
	assert.NotEqual(t, rl.Remaining, 0)

	assert.NotNil(t, res)
	assert.NotEqual(t, res.Total, 0)
	assert.NotEqual(t, res.TotalPages, 0)
	assert.NotEmpty(t, res.Results)
}

func TestClient_SearchUsers(t *testing.T) {
	c, err := unsplash.New(unsplash.WithHttpClient(httpClient))
	require.Nil(t, err)

	res, rl, err := c.SearchUsers(context.Background(), unsplash.SearchUsersOptions{Query: "car"})
	require.Nil(t, err)
	assert.NotNil(t, rl)
	assert.True(t, rl.Limit >= minLimits)
	assert.NotEqual(t, rl.Remaining, 0)

	assert.NotNil(t, res)
	assert.NotEqual(t, res.Total, 0)
	assert.NotEqual(t, res.TotalPages, 0)
	assert.NotEmpty(t, res.Results)
}
