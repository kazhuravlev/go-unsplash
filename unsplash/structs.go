package unsplash

type Exif struct {
	Make         string `json:"make"`
	Model        string `json:"model"`
	ExposureTime string `json:"exposure_time"`
	Aperture     string `json:"aperture"`
	FocalLength  string `json:"focal_length"`
	Iso          int    `json:"iso"`
}

type Position struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Location struct {
	Title    string   `json:"title"`
	Name     string   `json:"name"`
	City     string   `json:"city"`
	Country  string   `json:"country"`
	Position Position `json:"position"`
}

type CurrentUserCollection struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	PublishedAt string      `json:"published_at"`
	UpdatedAt   string      `json:"updated_at"`
	Curated     bool        `json:"curated"`
	CoverPhoto  interface{} `json:"cover_photo"`
	User        interface{} `json:"user"`
}

type Urls struct {
	Raw     string `json:"raw"`
	Full    string `json:"full"`
	Regular string `json:"regular"`
	Small   string `json:"small"`
	Thumb   string `json:"thumb"`
}

type PhotoLinks struct {
	Self             string `json:"self"`
	HTML             string `json:"html"`
	Download         string `json:"download"`
	DownloadLocation string `json:"download_location"`
}

type UserLinks struct {
	Self      string `json:"self"`
	HTML      string `json:"html"`
	Photos    string `json:"photos"`
	Likes     string `json:"likes"`
	Portfolio string `json:"portfolio"`
	Following string `json:"following"`
	Followers string `json:"followers"`
}

type User struct {
	ID                string    `json:"id"`
	UpdatedAt         string    `json:"updated_at"`
	Username          string    `json:"username"`
	Name              string    `json:"name"`
	PortfolioURL      string    `json:"portfolio_url"`
	Bio               string    `json:"bio"`
	Location          string    `json:"location"`
	TotalLikes        int       `json:"total_likes"`
	TotalPhotos       int       `json:"total_photos"`
	TotalCollections  int       `json:"total_collections"`
	Links             UserLinks `json:"links"`
	InstagramUsername string    `json:"instagram_username"`
	AcceptedTos       bool      `json:"accepted_tos"`
	ProfileImage      struct {
		Small  string `json:"small"`
		Medium string `json:"medium"`
		Large  string `json:"large"`
	} `json:"profile_image"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	TwitterUsername string `json:"twitter_username"`
}

type Photo struct {
	ID                     string                  `json:"id"`
	CreatedAt              string                  `json:"created_at"`
	UpdatedAt              string                  `json:"updated_at"`
	Width                  int                     `json:"width"`
	Height                 int                     `json:"height"`
	Color                  string                  `json:"color"`
	Description            string                  `json:"description"`
	Sponsored              bool                    `json:"sponsored"`
	SponsoredBy            interface{}             `json:"sponsored_by"`
	SponsoredImpressionsID interface{}             `json:"sponsored_impressions_id"`
	Downloads              int                     `json:"downloads"`
	Likes                  int                     `json:"likes"`
	LikedByUser            bool                    `json:"liked_by_user"`
	Exif                   Exif                    `json:"exif"`
	Location               Location                `json:"location"`
	CurrentUserCollections []CurrentUserCollection `json:"current_user_collections"`
	Urls                   Urls                    `json:"urls"`
	Links                  PhotoLinks              `json:"links"`
	User                   User                    `json:"user"`
	Categories             []string                `json:"categories"`
	Views                  int                     `json:"views"`
	Slug                   string                  `json:"slug"`
}

type Stat struct {
	Total      int `json:"total"`
	Historical struct {
		Change     int    `json:"change"`
		Resolution string `json:"resolution"`
		Quantity   int    `json:"quantity"`
		Values     []struct {
			Date  string `json:"date"`
			Value int    `json:"value"`
		} `json:"values"`
	} `json:"historical"`
}

type PhotoStatistics struct {
	ID        string `json:"id"`
	Downloads Stat   `json:"downloads"`
	Views     Stat   `json:"views"`
	Likes     Stat   `json:"likes"`
}

type PhotoDownload struct {
	URL string `json:"url"`
}

type Tag struct {
	Title string `json:"title"`
}

type SearchPhoto struct {
	Photo
	Tags      []Tag `json:"tags"`
	PhotoTags []Tag `json:"photo_tags"`
}

type SearchResult struct {
	Total      int           `json:"total"`
	TotalPages int           `json:"total_pages"`
	Results    []SearchPhoto `json:"results"`
}

type CollectionLinks struct {
	Self    string `json:"self"`
	HTML    string `json:"html"`
	Photos  string `json:"photos"`
	Related string `json:"related"`
}

type PreviewPhoto struct {
	ID   string `json:"id"`
	Urls Urls   `json:"urls"`
}

type Collection struct {
	ID            int             `json:"id"`
	Title         string          `json:"title"`
	Description   string          `json:"description"`
	PublishedAt   string          `json:"published_at"`
	UpdatedAt     string          `json:"updated_at"`
	Curated       bool            `json:"curated"`
	Featured      bool            `json:"featured"`
	TotalPhotos   int             `json:"total_photos"`
	Private       bool            `json:"private"`
	ShareKey      string          `json:"share_key"`
	Tags          []Tag           `json:"tags"`
	CoverPhoto    Photo           `json:"cover_photo"`
	PreviewPhotos []PreviewPhoto  `json:"preview_photos"`
	User          User            `json:"user"`
	Links         CollectionLinks `json:"links"`
}

type CollectionSearchResult struct {
	Total      int          `json:"total"`
	TotalPages int          `json:"total_pages"`
	Results    []Collection `json:"results"`
}

type UsersSearchResult struct {
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	Results    []User `json:"results"`
}
