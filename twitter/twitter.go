package twitter

import (
	"github.com/dghubble/sling"
	"net/http"
)

const twitterAPI = "https://api.twitter.com/1.1/"

// Client is a Twitter client for making Twitter API requests.
type Client struct {
	sling *sling.Sling
	// Twitter API Services
	Statuses  *StatusService
	Timelines *TimelineService
	Users     *UserService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(twitterAPI)
	return &Client{
		sling:     base,
		Statuses:  NewStatusService(base.New()),
		Timelines: NewTimelineService(base.New()),
		Users:     NewUserService(base.New()),
	}
}

// Bool returns a new pointer to the given bool value.
func Bool(v bool) *bool {
	ptr := new(bool)
	*ptr = v
	return ptr
}
