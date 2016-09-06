package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

// FriendshipService provides methods for accessing Twitter friendship endpoints.
type FriendshipService struct {
	sling *sling.Sling
}

// Creates a new friendship service
func newFriendshipService(sling *sling.Sling) *FriendshipService {
	return &FriendshipService{
		sling: sling.Path("friendships/"),
	}
}

// FriendshipLookupStatus is The relationship status between the authenticated user and the target
type FriendshipLookupStatus struct {
	Name        string   `json:"name"`
	ScreenName  string   `json:"screen_name"`
	ID          int64    `json:"id"`
	IDStr       string   `json:"id_str"`
	Connections []string `json:"connections"`
}

// FriendshipLookupParams are Basic parameters for friendship requests
type FriendshipLookupParams struct {
	UserID     string `url:"user_id,omitempty"`
	ScreenName string `url:"screen_name,omitempty"`
}

// Lookup returns the relationships of the authenticating user to target user
func (s *FriendshipService) Lookup(params *FriendshipLookupParams) (*[]FriendshipLookupStatus, *http.Response, error) {
	friendships := new([]FriendshipLookupStatus)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("lookup.json").QueryStruct(params).Receive(friendships, apiError)
	return friendships, resp, relevantError(err, *apiError)
}

// FriendshipShowResult is the result from the Friendship show function
type FriendshipShowResult struct {
	Relationship FriendshipRelationship `json:"relationship"`
}

// FriendshipRelationship is the underlying relationship of the show function
type FriendshipRelationship struct {
	Target FriendshipRelationshipTarget `json:"target"`
	Source FriendshipRelationshipSource `json:"source"`
}

// FriendshipRelationshipTarget is the target's attributes from the show function
type FriendshipRelationshipTarget struct {
	IDStr      string `json:"id_str"`
	ID         int64  `json:"id"`
	ScreenName string `json:"screen_name"`
	Following  bool   `json:"following"`
	FollowedBy bool   `json:"followed_by"`
}

// FriendshipRelationshipSource is the source's attributes from the show function
type FriendshipRelationshipSource struct {
	CanDM                bool   `json:"can_dm"`
	Blocking             bool   `json:"blocking"`
	Muting               bool   `json:"muting"`
	IDStr                string `json:"id_str"`
	AllReplies           bool   `json:"all_replies"`
	WantRetweets         bool   `json:"want_retweets"`
	ID                   int64  `json:"id"`
	MarkedSpam           bool   `json:"marked_spam"`
	ScreenName           string `json:"screen_name"`
	Following            bool   `json:"following"`
	FollowedBy           bool   `json:"followed_by"`
	NotificationsEnabled bool   `json:"notifications_enabled"`
}

// FriendshipShowParams are the parameters given to the show function
type FriendshipShowParams struct {
	SourceScreenName string `url:"source_screen_name,omitempty"`
	SourceID         string `url:"source_id,omitempty"`
	TargetScreenName string `url:"target_screen_name,omitempty"`
	TargetID         string `url:"target_id,omitempty"`
}

// Show returns the relationship between any two specified users
func (s *FriendshipService) Show(params *FriendshipShowParams) (*FriendshipShowResult, *http.Response, error) {
	friendships := new(FriendshipShowResult)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("show.json").QueryStruct(params).Receive(friendships, apiError)
	return friendships, resp, relevantError(err, *apiError)
}

// Destroy unfollows a user
func (s *FriendshipService) Destroy(params *FriendshipLookupParams) (*User, *http.Response, error) {
	friendships := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("destroy.json").QueryStruct(params).Receive(friendships, apiError)
	return friendships, resp, relevantError(err, *apiError)
}

// Create follows a user
func (s *FriendshipService) Create(params *FriendshipLookupParams) (*User, *http.Response, error) {
	friendships := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("create.json").QueryStruct(params).Receive(friendships, apiError)
	return friendships, resp, relevantError(err, *apiError)
}
