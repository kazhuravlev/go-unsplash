package unsplash_test

import (
	"context"
	"fmt"
	"github.com/kazhuravlev/go-unsplash/unsplash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
	"testing"
)

var (
	accessKey   string
	secretKey   string
	accessToken string

	httpClient *http.Client
)

func TestMain(m *testing.M) {
	accessKey = os.Getenv("TEST_ACCESS_KEY")
	if accessKey == "" {
		log.Fatalf("Please set 'TEST_ACCESS_KEY' env variable")
	}

	secretKey = os.Getenv("TEST_SECRET_KEY")
	if secretKey == "" {
		log.Fatalf("Please set 'TEST_SECRET_KEY' env variable")
	}

	accessToken = os.Getenv("TEST_ACCESS_TOKEN")
	if accessToken == "" {
		log.Fatal("Please set 'TEST_ACCESS_TOKEN' env varible")
	}

	source := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken, TokenType: "bearer"},
	)

	httpClient = oauth2.NewClient(context.Background(), source)

	os.Exit(m.Run())
}

func TestGetRandomPhotos(t *testing.T) {
	c, err := unsplash.New(unsplash.WithHttpClient(httpClient))
	require.Nil(t, err)

	photos, err := c.GetRandomPhotos(context.Background(), unsplash.GetRandomPhotosOptions{Orientation: unsplash.OrientationLandscape})
	require.Nil(t, err)

	assert.NotNil(t, photos)
	assert.Len(t, photos, 1)

	fmt.Println(photos[0].Urls.Full)
}
