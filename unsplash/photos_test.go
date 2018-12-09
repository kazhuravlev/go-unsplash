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
