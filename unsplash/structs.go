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
}

type User struct {
	ID               string    `json:"id"`
	UpdatedAt        string    `json:"updated_at"`
	Username         string    `json:"username"`
	Name             string    `json:"name"`
	PortfolioURL     string    `json:"portfolio_url"`
	Bio              string    `json:"bio"`
	Location         string    `json:"location"`
	TotalLikes       int       `json:"total_likes"`
	TotalPhotos      int       `json:"total_photos"`
	TotalCollections int       `json:"total_collections"`
	Links            UserLinks `json:"links"`
}

type Photo struct {
	ID                     string                  `json:"id"`
	CreatedAt              string                  `json:"created_at"`
	UpdatedAt              string                  `json:"updated_at"`
	Width                  int                     `json:"width"`
	Height                 int                     `json:"height"`
	Color                  string                  `json:"color"`
	Downloads              int                     `json:"downloads"`
	Likes                  int                     `json:"likes"`
	LikedByUser            bool                    `json:"liked_by_user"`
	Description            string                  `json:"description"`
	Exif                   Exif                    `json:"exif"`
	Location               Location                `json:"location"`
	CurrentUserCollections []CurrentUserCollection `json:"current_user_collections"`
	Urls                   Urls                    `json:"urls"`
	Links                  PhotoLinks              `json:"links"`
	User                   User                    `json:"user"`
}
