package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/kazhuravlev/go-unsplash/unsplash"
	"golang.org/x/oauth2"
	"log"
)

var (
	scopes = []unsplash.Scope{
		unsplash.ScopePublic,
		unsplash.ScopeReadUser,
		unsplash.ScopeWriteUser,
		unsplash.ScopeReadPhotos,
		unsplash.ScopeWritePhotos,
		unsplash.ScopeWriteLikes,
		unsplash.ScopeWriteFollowers,
		unsplash.ScopeReadCollections,
		unsplash.ScopeWriteCollections,
	}
	redirectURL = "urn:ietf:wg:oauth:2.0:oob"
)

func main() {
	var accessKey, secretKey, code string
	flag.StringVar(&accessKey, "accessKey", "", "-accessKey=my_key")
	flag.StringVar(&secretKey, "secretKey", "", "-secretKey=my_secret_key")
	flag.StringVar(&code, "code", "", "-code=code")
	flag.Parse()

	if accessKey == "" || secretKey == "" {
		log.Fatal("Please specify 'accessKey' and 'secretKey'")
	}

	sScopes := unsplash.Scopes(scopes...)

	conf := &oauth2.Config{
		ClientID:     accessKey,
		ClientSecret: secretKey,
		Scopes:       sScopes,
		RedirectURL:  redirectURL,
		Endpoint:     unsplash.Oauth2Endpoint,
	}

	if code == "" {
		fmt.Println("Go to URL and get 'code' string")

		u := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)

		fmt.Println(u)
	} else {
		token, err := conf.Exchange(context.Background(), code)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Your 'accessToken':")
		fmt.Println(token.AccessToken)
	}
}
