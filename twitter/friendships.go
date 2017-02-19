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

// FriendshipShowParams are paramenters for FriendshipService.Show
type FriendshipShowParams struct {
	SourceID         int64  `url:"source_id,omitempty"`
	SourceScreenName string `url:"source_screen_name,omitempty"`
	TargetID         int64  `url:"target_id,omitempty"`
	TargetScreenName string `url:"target_screen_name,omitempty"`
}

// Show returns the relationship between two arbitrary users.
// Requires a user auth or an app context.
// https://dev.twitter.com/rest/reference/get/friendships/show
func (s *FriendshipService) Show(params *FriendshipShowParams) (*Relationship, *http.Response, error) {
	response := new(RelationshipResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("show.json").QueryStruct(params).Receive(response, apiError)
	return response.Relationship, resp, relevantError(err, *apiError)
}

// RelationshipResponse contains a relationship.
type RelationshipResponse struct {
	Relationship *Relationship `json:"relationship"`
}

// Relationship represents the relation between a source user and target user.
type Relationship struct {
	Source RelationshipSource `json:"source"`
	Target RelationshipTarget `json:"target"`
}

// RelationshipSource represents the source user.
type RelationshipSource struct {
	ID                   int64  `json:"id"`
	IDStr                string `json:"id_str"`
	ScreenName           string `json:"screen_name"`
	Following            bool   `json:"following"`
	FollowedBy           bool   `json:"followed_by"`
	CanDM                bool   `json:"can_dm"`
	Blocking             bool   `json:"blocking"`
	Muting               bool   `json:"muting"`
	AllReplies           bool   `json:"all_replies"`
	WantRetweets         bool   `json:"want_retweets"`
	MarkedSpam           bool   `json:"marked_spam"`
	NotificationsEnabled bool   `json:"notifications_enabled"`
}

// RelationshipTarget represents the target user.
type RelationshipTarget struct {
	ID         int64  `json:"id"`
	IDStr      string `json:"id_str"`
	ScreenName string `json:"screen_name"`
	Following  bool   `json:"following"`
	FollowedBy bool   `json:"followed_by"`
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

// FriendshipShowParams are paramenters for FriendshipService.Show
type FriendshipShowParams struct {
	SourceID         int64  `url:"source_id,omitempty"`
	SourceScreenName string `url:"source_screen_name,omitempty"`
	TargetID         int64  `url:"target_id,omitempty"`
	TargetScreenName string `url:"target_screen_name,omitempty"`
}

// Show returns the relationship between two arbitrary users.
// Requires a user auth or an app context.
// https://dev.twitter.com/rest/reference/get/friendships/show
func (s *FriendshipService) Show(params *FriendshipShowParams) (*Relationship, *http.Response, error) {
	relationship := new(Relationship)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("show.json").QueryStruct(params).Receive(relationship, apiError)
	return relationship, resp, relevantError(err, *apiError)
}

// Relationship represents the relationship between two arbitrary users.
type Relationship struct {
	*RelationshipResponse `json:"relationship"`
}

// RelationshipResponse is struct for relationship API response.
type RelationshipResponse struct {
	Source *RelationshipSource `json:"source"`
	Target *RelationshipTarget `json:"target"`
}

// RelationshipSource represents the source user.
type RelationshipSource struct {
	ID                   int64  `json:"id"`
	IDStr                string `json:"id_str"`
	ScreenName           string `json:"screen_name"`
	Following            bool   `json:"following"`
	FollowedBy           bool   `json:"followed_by"`
	CanDm                bool   `json:"can_dm"`
	Blocking             *bool  `json:"blocking"`
	Muting               *bool  `json:"muting"`
	AllReplies           *bool  `json:"all_replies"`
	WantRetweets         *bool  `json:"want_retweets"`
	MarkedSpam           *bool  `json:"marked_spam"`
	NotificationsEnabled *bool  `json:"notifications_enabled"`
}

// RelationshipTarget represents the target user.
type RelationshipTarget struct {
	IDStr       string `json:"id_str"`
	ID          int64  `json:"id"`
	ScreenName  string `json:"screen_name"`
	Following   bool   `json:"following"`
	FollowingBy bool   `json:"followed_by"`
}
