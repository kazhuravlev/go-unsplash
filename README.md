# Go client for unsplash.com API

## Installation

```bash
go get github.com/kazhuravlev/go-unsplash 
```

## Run tests

To run tests you must [register application](https://unsplash.com/oauth/applications/new).

After register application copy `Access Key` and `Secret Key` and then:

```bash
go run tools/main.go -accessKey=<Access Key> -secretKey=<Secret Key>
```

Go to URL and authorize application to required permissions. After accept 
request you get `Authorization code` string. Copy it and:

```bash
go run tools/main.go -accessKey=<Access Key> -secretKey=<Secret Key> -code=<Authorization code> 
```

You must see `Access Token` in terminal output. Copy it.

To run all tests with given credentials just type:

```bash
export TEST_ACCESS_KEY=<Access Key>; TEST_SECRET_KEY=<Secret Key>; TEST_ACCESS_TOKEN=<Access Token> go test -v ./...
```

## Usage

```go
package main

import (
	"context"
	"fmt"
	"github.com/kazhuravlev/go-unsplash/unsplash"
	"golang.org/x/oauth2"
	"log"
)

func main() {
	accessToken := "<Access Key>"
	source := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken, TokenType: "Client-ID"},
	)

	httpClient := oauth2.NewClient(context.Background(), source)

	client, err := unsplash.New(unsplash.WithHttpClient(httpClient))
	if err != nil {
		log.Fatal(err)
	}

	photos, err := client.GetRandomPhotos(context.Background(), unsplash.GetRandomPhotosOptions{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(photos[0].Urls.Full)
}
```