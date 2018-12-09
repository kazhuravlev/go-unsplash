package unsplash

import "golang.org/x/oauth2"

var (
	Oauth2Endpoint = oauth2.Endpoint{
		AuthURL:  "https://unsplash.com/oauth/authorize",
		TokenURL: "https://unsplash.com/oauth/token",
	}
)

type Scope string

const (
	// ScopePublic Default. Read public data.
	ScopePublic Scope = "public"
	// ScopeRead_user Access user’s private data.
	ScopeReadUser Scope = "read_user"
	// ScopeWrite_user Update the user’s profile.
	ScopeWriteUser Scope = "write_user"
	// ScopeRead_photos Read private data from the user’s photos.
	ScopeReadPhotos Scope = "read_photos"
	// ScopeWrite_photos Update photos on the user’s behalf.
	ScopeWritePhotos Scope = "write_photos"
	// ScopeWrite_likes Like or unlike a photo on the user’s behalf.
	ScopeWriteLikes Scope = "write_likes"
	// ScopeWrite_followers Follow or unfollow a user on the user’s behalf.
	ScopeWriteFollowers Scope = "write_followers"
	// ScopeRead_collections View a user’s private collections.
	ScopeReadCollections Scope = "read_collections"
	// ScopeWrite_collections Create and update a user’s collections.
	ScopeWriteCollections Scope = "write_collections"
)

// Scopes convert slice of Scope to slice of string for oatuh2.Config.Scopes
func Scopes(scopes ...Scope) []string {
	sScopes := make([]string, len(scopes))
	for i, scope := range scopes {
		sScopes[i] = string(scope)
	}

	return sScopes
}
