package unsplash_test

import (
	"context"
	"fmt"
	"github.com/kazhuravlev/go-unsplash/unsplash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_GetRandomPhotos(t *testing.T) {
	c, err := unsplash.New(unsplash.WithHttpClient(httpClient))
	require.Nil(t, err)

	n := 30

	photos, err := c.GetRandomPhotos(context.Background(), unsplash.GetRandomPhotosOptions{Orientation: unsplash.OrientationLandscape, Count: n})
	require.Nil(t, err)

	assert.NotNil(t, photos)
	assert.Len(t, photos, n)

	fmt.Println(photos[0].Urls.Full)
}

func TestClient_GetPhotos(t *testing.T) {
	c, err := unsplash.New(unsplash.WithHttpClient(httpClient))
	require.Nil(t, err)

	n := 30
	photos, err := c.GetPhotos(context.Background(), unsplash.GetPhotosOptions{PerPage: n})
	require.Nil(t, err)

	assert.NotNil(t, photos)
	assert.Len(t, photos, n)

	fmt.Println(photos[0].Urls.Full)
}

func TestClient_GetCuratedPhotos(t *testing.T) {
	c, err := unsplash.New(unsplash.WithHttpClient(httpClient))
	require.Nil(t, err)

	n := 30
	photos, err := c.GetCuratedPhotos(context.Background(), unsplash.GetPhotosOptions{PerPage: n})
	require.Nil(t, err)

	assert.NotNil(t, photos)
	assert.Len(t, photos, n)

	fmt.Println(photos[0].Urls.Full)
}

func TestClient_GetPhoto(t *testing.T) {
	c, err := unsplash.New(unsplash.WithHttpClient(httpClient))
	require.Nil(t, err)

	id := "Mg0W1N_yDv0"
	photo, err := c.GetPhoto(context.Background(), id)
	require.Nil(t, err)

	assert.NotNil(t, photo)
	assert.Equal(t, photo.ID, id)
}

func TestClient_GetPhotoStatistics(t *testing.T) {
	c, err := unsplash.New(unsplash.WithHttpClient(httpClient))
	require.Nil(t, err)

	id := "Mg0W1N_yDv0"
	photo, err := c.GetPhotoStatistics(context.Background(), unsplash.GetPhotoStatisticsOptions{ID: id})
	require.Nil(t, err)

	assert.NotNil(t, photo)
	assert.Equal(t, photo.ID, id)
}

func TestClient_GetPhotoDownload(t *testing.T) {
	c, err := unsplash.New(unsplash.WithHttpClient(httpClient))
	require.Nil(t, err)

	id := "Mg0W1N_yDv0"
	photo, err := c.GetPhotoDownload(context.Background(), id)
	require.Nil(t, err)

	assert.NotNil(t, photo)
	assert.Equal(t, photo.URL, "https://images.unsplash.com/photo-1536167038724-17be8c5e6876?ixlib=rb-1.2.1&q=85&fm=jpg&crop=entropy&cs=srgb")
}
