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
	DirectMessages *DirectMessageService
	Favorites      *FavoriteService
	Followers      *FollowerService
	Friends        *FriendService
	Friendships    *FriendshipService
	Lists          *ListsService
	RateLimits     *RateLimitService
	Search         *SearchService
	PremiumSearch  *PremiumSearchService
	Statuses       *StatusService
	Streams        *StreamService
	Timelines      *TimelineService
	Trends         *TrendsService
	Users          *UserService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(twitterAPI)
	return &Client{
		sling:          base,
		Accounts:       newAccountService(base.New()),
		DirectMessages: newDirectMessageService(base.New()),
		Favorites:      newFavoriteService(base.New()),
		Followers:      newFollowerService(base.New()),
		Friends:        newFriendService(base.New()),
		Friendships:    newFriendshipService(base.New()),
		Lists:          newListService(base.New()),
		RateLimits:     newRateLimitService(base.New()),
		Search:         newSearchService(base.New()),
		PremiumSearch:  newPremiumSearchService(base.New()),
		Statuses:       newStatusService(base.New()),
		Streams:        newStreamService(httpClient, base.New()),
		Timelines:      newTimelineService(base.New()),
		Trends:         newTrendsService(base.New()),
		Users:          newUserService(base.New()),
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
