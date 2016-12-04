package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

// FriendshipService provides methods for accessing Twitter friendship API
// endpoints.
type FriendshipService struct {
	sling *sling.Sling
}

// newFriendshipService returns a new FriendshipService.
func newFriendshipService(sling *sling.Sling) *FriendshipService {
	return &FriendshipService{
		sling: sling.Path("friendships/"),
	}
}

// FriendshipCreateParams are parameters for FriendshipService.Create
type FriendshipCreateParams struct {
	ScreenName string `url:"screen_name,omitempty"`
	UserID     int64  `url:"user_id,omitempty"`
	Follow     *bool  `url:"follow,omitempty"`
}

// Create creates a friendship to (i.e. follows) the specified user and
// returns the followed user.
// Requires a user auth context.
// https://dev.twitter.com/rest/reference/post/friendships/create
func (s *FriendshipService) Create(params *FriendshipCreateParams) (*User, *http.Response, error) {
	user := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("create.json").QueryStruct(params).Receive(user, apiError)
	return user, resp, relevantError(err, *apiError)
}

// FriendshipDestroyParams are paramenters for FriendshipService.Destroy
type FriendshipDestroyParams struct {
	ScreenName string `url:"screen_name,omitempty"`
	UserID     int64  `url:"user_id,omitempty"`
}

// Destroy destroys a friendship to (i.e. unfollows) the specified user and
// returns the unfollowed user.
// Requires a user auth context.
// https://dev.twitter.com/rest/reference/post/friendships/destroy
func (s *FriendshipService) Destroy(params *FriendshipDestroyParams) (*User, *http.Response, error) {
	user := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("destroy.json").QueryStruct(params).Receive(user, apiError)
	return user, resp, relevantError(err, *apiError)
}
