package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

type FriendshipService struct {
	sling *sling.Sling
}

func newFriendshipService(sling *sling.Sling) *FriendshipService {
	return &FriendshipService{
		sling: sling.Path("friendships/"),
	}
}

type FriendshipLookupStatus struct {
	Name        string   `json:"name"`
	ScreenName  string   `json:"screen_name"`
	Id          int64    `json:"id"`
	IdStr       string   `json:"id_str"`
	Connections []string `json:"connections"`
}

type FriendshipLookupParams struct {
	UserID     string `url:"user_id,omitempty"`
	ScreenName string `url:"screen_name,omitempty"`
}

func (s *FriendshipService) Lookup(params *FriendshipLookupParams) (*[]FriendshipLookupStatus, *http.Response, error) {
	friendships := new([]FriendshipLookupStatus)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("lookup.json").QueryStruct(params).Receive(friendships, apiError)
	return friendships, resp, relevantError(err, *apiError)
}

type FriendshipShowResult struct {
	Relationship FriendshipRelationship `json:"relationship"`
}

type FriendshipRelationship struct {
	Target FriendshipRelationshipTarget `json:"target"`
	Source FriendshipRelationshipSource `json:"source"`
}

type FriendshipRelationshipTarget struct {
	IdStr      string `json:"id_str"`
	Id         int64  `json:"id"`
	ScreenName string `json:"screen_name"`
	Following  bool   `json:"following"`
	FollowedBy bool   `json:"followed_by"`
}

type FriendshipRelationshipSource struct {
	CanDM                bool   `json:"can_dm"`
	Blocking             bool   `json:"blocking"`
	Muting               bool   `json:"muting"`
	IdStr                string `json:"id_str"`
	AllReplies           bool   `json:"all_replies"`
	WantRetweets         bool   `json:"want_retweets"`
	Id                   int64  `json:"id"`
	MarkedSpam           bool   `json:"marked_spam"`
	ScreenName           string `json:"screen_name"`
	Following            bool   `json:"following"`
	FollowedBy           bool   `json:"followed_by"`
	NotificationsEnabled bool   `json:"notifications_enabled"`
}

type FriendshipShowParams struct {
	SourceScreenName string `url:"source_screen_name,omitempty"`
	SourceId         string `url:"source_id,omitempty"`
	TargetScreenName string `url:"target_screen_name,omitempty"`
	TargetId         string `url:"target_id,omitempty"`
}

func (s *FriendshipService) Show(params *FriendshipShowParams) (*FriendshipShowResult, *http.Response, error) {
	friendships := new(FriendshipShowResult)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("show.json").QueryStruct(params).Receive(friendships, apiError)
	return friendships, resp, relevantError(err, *apiError)
}

type FriendshipDestroyResult struct {
	Name string `json:"name"`
	Id   int64  `json:"id"`
}

func (s *FriendshipService) Destroy(params *FriendshipLookupParams) (*FriendshipDestroyResult, *http.Response, error) {
	friendships := new(FriendshipDestroyResult)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("destroy.json").QueryStruct(params).Receive(friendships, apiError)
	return friendships, resp, relevantError(err, *apiError)
}
