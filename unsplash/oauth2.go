package unsplash

import "golang.org/x/oauth2"

var (
	Oauth2Endpoint = oauth2.Endpoint{
		AuthURL:  "https://unsplash.com/oauth/authorize",
		TokenURL: "https://unsplash.com/oauth/token",
	}
)
