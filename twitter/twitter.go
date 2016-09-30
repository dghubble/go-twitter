package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

const twitterAPI = "https://api.twitter.com/1.1/"

// Client is a Twitter client for making Twitter API requests.
type Client struct {
	sling *sling.Sling
	// Twitter API Services
	Accounts       *AccountService
	Statuses       *StatusService
	Timelines      *TimelineService
	Users          *UserService
	Favorites      *FavoriteService
	Followers      *FollowerService
	DirectMessages *DirectMessageService
	Search         *SearchService
	Streams        *StreamService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(twitterAPI)
	return &Client{
		sling:          base,
		Accounts:       newAccountService(base.New()),
		Statuses:       newStatusService(base.New()),
		Timelines:      newTimelineService(base.New()),
		Users:          newUserService(base.New()),
		Favorites:      newFavoriteService(base.New()),
		Followers:      newFollowerService(base.New()),
		DirectMessages: newDirectMessageService(base.New()),
		Search:         newSearchService(base.New()),
		Streams:        newStreamService(httpClient, base.New()),
	}
}

// Bool returns a new pointer to the given bool value.
func Bool(v bool) *bool {
	ptr := new(bool)
	*ptr = v
	return ptr
}

// Float returns a new pointer to the given float64 value.
func Float(v float64) *float64 {
	ptr := new(float64)
	*ptr = v
	return ptr
}
