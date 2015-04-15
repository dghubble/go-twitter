package twitter

import (
	"github.com/dghubble/sling"
	"net/http"
)

const twitterApi = "https://api.twitter.com/1.1/"

// API Client communicates with the Twitter API services.
type Client struct {
	sling *sling.Sling
	// Different Twitter API Services
	Users *UserService
}

func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(twitterApi)
	return &Client{
		sling: base,
		Users: NewUserService(base.New()),
	}
}
