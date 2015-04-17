package twitter

import (
	"github.com/dghubble/sling"
	"net/http"
)

// https://dev.twitter.com/overview/api/users
type User struct {
	ContributorsEnabled            bool          `json:"contributors_enabled"`
	CreatedAt                      string        `json:"created_at"`
	DefaultProfile                 bool          `json:"default_profile"`
	DefaultProfileImage            bool          `json:"default_profile_image"`
	Description                    string        `json:"description"`
	Entities                       *UserEntities `json:"entities"`
	FavouritesCount                int           `json:"favourites_count"`
	FollowRequestSent              bool          `json:"follow_request_sent"`
	Following                      bool          `json:"following"`
	FollowersCount                 int           `json:"followers_count"`
	FriendsCount                   int           `json:"friends_count"`
	GeoEnabled                     bool          `json:"geo_enabled"`
	Id                             int64         `json:"id"`
	IdStr                          string        `json:"id_str"`
	IsTranslator                   bool          `json:"id_translator"`
	Lang                           string        `json:"lang"`
	ListedCount                    int           `json:"listed_count"`
	Location                       string        `json:"location"`
	Name                           string        `json:"name"`
	Notifications                  bool          `json:"notifications"`
	ProfileBackgroundColor         string        `json:"profile_background_color"`
	ProfileBackgroundImageUrl      string        `json:"profile_background_image_url"`
	ProfileBackgroundImageUrlHttps string        `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool          `json:"profile_background_tile"`
	ProfileBannerUrl               string        `json:"profile_banner_url"`
	ProfileImageUrl                string        `json:"profile_image_url"`
	ProfileImageUrlHttps           string        `json:"profile_image_url_https"`
	ProfileLinkColor               string        `json:"profile_link_color"`
	ProfileSidebarBorderColor      string        `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string        `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string        `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool          `json:"profile_use_background_image"`
	Protected                      bool          `json:"protected"`
	ScreenName                     string        `json:"screen_name"`
	ShowAllInlineMedia             bool          `json:"show_all_inline_media"`
	Status                         *Tweet        `json:"status"`
	StatusesCount                  int           `json:"statuses_count"`
	Timezone                       string        `json:"time_zone"`
	Url                            string        `json:"url"`
	UtcOffset                      int           `json:"utc_offset"`
	Verified                       bool          `json:"verified"`
	WithheldInCountries            string        `json:"withheld_in_countries"`
	WithholdScope                  string        `json:"withheld_scope"`
}

type UserService struct {
	sling *sling.Sling
}

func NewUserService(sling *sling.Sling) *UserService {
	return &UserService{
		sling: sling.Path("users/"),
	}
}

// https://dev.twitter.com/rest/reference/get/users/show
type UserShowParams struct {
	UserId          int64  `url:"user_id,omitempty"`
	ScreenName      string `url:"screen_name,omitempty"`
	IncludeEntities *bool  `url:"include_entities,omitempty"` // whether 'status' should include entities
}

// Show returns the specified User.
// https://dev.twitter.com/rest/reference/get/users/show
func (s *UserService) Show(params *UserShowParams) (*User, *http.Response, error) {
	user := new(User)
	resp, err := s.sling.New().Get("show.json").QueryStruct(params).Receive(user)
	return user, resp, err
}

// https://dev.twitter.com/rest/reference/get/users/lookup
type UserLookupParams struct {
	UserId          []int64  `url:"user_id,omitempty"`
	ScreenName      []string `url:"screen_name,omitempty"`
	IncludeEntities *bool    `url:"include_entities,omitempty"` // whether 'status' should include entities
}

// Lookup returns the specified slice of Users.
// https://dev.twitter.com/rest/reference/get/users/lookup
func (s *UserService) Lookup(params *UserLookupParams) ([]User, *http.Response, error) {
	users := new([]User)
	resp, err := s.sling.New().Get("lookup.json").QueryStruct(params).Receive(users)
	return *users, resp, err
}

// https://dev.twitter.com/rest/reference/get/users/search
type UserSearchParams struct {
	Query           string `url:"q,omitempty"`
	Page            int    `url:"page,omitempty"` // 1-based page number
	Count           int    `url:"count,omitempty"`
	IncludeEntities *bool  `url:"include_entities,omitempty"` // whether 'status' should include entities
}

// Search queries public user accounts. User auth only.
// https://dev.twitter.com/rest/reference/get/users/search
func (s *UserService) Search(query string, params *UserSearchParams) ([]User, *http.Response, error) {
	params.Query = query
	users := new([]User)
	resp, err := s.sling.New().Get("search.json").QueryStruct(params).Receive(users)
	return *users, resp, err
}
