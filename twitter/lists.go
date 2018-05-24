package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

// List represents a twitter list
type List struct {
	Slug            string `json:"slug"`
	Name            string `json:"name"`
	CreatedAt       string `json:"created_at"`
	URI             string `json:"uri"`
	SubscriberCount int    `json:"subscriber_count"`
	IDStr           string `json:"id_str"`
	MemberCount     int    `json:"member_count"`
	Mode            string `json:"mode"`
	ID              int    `json:"id"`
	FullName        string `json:"full_name"`
	Description     string `json:"description"`
	User            `json:"user"`
}

// Members represents the members of a twitter list
type Members struct {
	Users             []User `json:"users"`
	NextCursor        int64  `json:"next_cursor"`
	NextCursorStr     string `json:"next_cursor_str"`
	PreviousCursor    int64  `json:"previous_cursor"`
	PreviousCursorStr string `json:"previous_cursor_str"`
}

// Memberships represent the memberships of a twitter list
type Memberships struct {
	Lists             []List `json:"lists"`
	NextCursor        int64  `json:"next_cursor"`
	NextCursorStr     string `json:"next_cursor_str"`
	PreviousCursor    int64  `json:"previous_cursor"`
	PreviousCursorStr string `json:"previous_cursor_str"`
}

// Ownerships represent the ownership of a twitter list
type Ownerships Memberships

// ListsService provides methods for accessing Twitter list endpoints.
type ListsService struct {
	sling *sling.Sling
}

// newListsService returns a new ListsService.
func newListsService(sling *sling.Sling) *ListsService {
	return &ListsService{
		sling: sling.Path("lists/"),
	}
}

// ListsListParams are the parameters for ListsService.List
type ListsListParams struct {
	UserID     int64  `url:"user_id,omitempty"`
	ScreenName string `url:"screen_name,omitempty"`
	Reverse    bool   `url:"reverse,omitempty"`
}

// List returns all lists the authenticating or specified user subscribes to, including their own.
// The user is specified using the user_id or screen_name parameters.
// If no user is given, the authenticating user is used.
// https://developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/get-lists-list
func (s *ListsService) List(params *ListsListParams) (*[]List, *http.Response, error) {
	list := new([]List)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("list.json").QueryStruct(params).Receive(list, apiError)
	return list, resp, relevantError(err, *apiError)
}

// ListsMembersParams are the parameters for ListsService.Members
type ListsMembersParams struct {
	ListID          int64  `url:"list_id,omitempty"`
	Slug            string `url:"slug,omitempty"`
	OwnerScreenName string `url:"owner_screen_name,omitempty"`
	OwnerID         int64  `url:"owner_id,omitempty"`
	Count           int    `url:"count,omitempty"`
	Cursor          int64  `url:"cursor,omitempty"`
	IncludeEntities *bool  `url:"include_entities,omitempty"`
	SkipStatus      *bool  `url:"skip_status,omitempty"`
}

// Members returns the members of the specified list.
// Private list members will only be shown if the authenticated user owns the specified list.
// https://developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/get-lists-members
func (s *ListsService) Members(params *ListsMembersParams) (*Members, *http.Response, error) {
	members := new(Members)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("members.json").QueryStruct(params).Receive(members, apiError)
	return members, resp, relevantError(err, *apiError)
}

// ListsMembersShowParams are the parameters for ListsService.MembersShow
type ListsMembersShowParams struct {
	ListID          int64  `url:"list_id,omitempty"`
	Slug            string `url:"slug,omitempty"`
	UserID          int64  `url:"user_id,omitempty"`
	ScreenName      string `url:"screen_name,omitempty"`
	OwnerScreenName string `url:"owner_screen_name,omitempty"`
	OwnerID         int64  `url:"slug,omitempty"`
	IncludeEntities *bool  `url:"include_entities,omitempty"`
	SkipStatus      *bool  `url:"skip_status,omitempty"`
}

// MembersShow checks if the specified user is a member of the specified list.
// https://developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/get-lists-members-show
func (s *ListsService) MembersShow(params *ListsMembersShowParams) (*User, *http.Response, error) {
	user := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("members/show.json").QueryStruct(params).Receive(user, apiError)
	return user, resp, relevantError(err, *apiError)
}

// ListsMembershipsParams are the parameters for ListsService.Memberships
type ListsMembershipsParams struct {
	UserID             int64  `url:"user_id,omitempty"`
	ScreenName         string `url:"screen_name,omitempty"`
	Count              int    `url:"count,omitempty"`
	Cursor             int64  `url:"cursor,omitempty"`
	FilterToOwnedLists *bool  `url:"filter_to_owned_lists,omitempty"`
}

// Memberships returns the lists the specified user has been added to.
// If user_id or screen_name are not provided the memberships for the authenticating user are returned.
// https://developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/get-lists-memberships
func (s *ListsService) Memberships(params *ListsMembershipsParams) (*Memberships, *http.Response, error) {
	memberships := new(Memberships)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("memberships.json").QueryStruct(params).Receive(memberships, apiError)
	return memberships, resp, relevantError(err, *apiError)
}

// ListsOwnershipsParams are the parameters for ListsService.Ownerships
type ListsOwnershipsParams struct {
	UserID             int64  `url:"user_id,omitempty"`
	ScreenName         string `url:"screen_name,omitempty"`
	Count              int    `url:"count,omitempty"`
	Cursor             int64  `url:"cursor,omitempty"`
	FilterToOwnedLists *bool  `url:"filter_to_owned_lists,omitempty"`
}

/*
TODO: Implement POST methods
POST lists/members/create
POST lists/members/create_all
POST lists/members/destroy
POST lists/members/destroy_all
POST lists/subscribers/create
POST lists/subscribers/destroy
POST lists/update
POST lists/create
POST lists/destroy

// ListsMembersCreateParams are the parameters for ListsService.MembersCreate
type ListsMembersCreateParams struct {
	ListID          int64  `url:"list_id,omitempty"`
	Slug            string `url:"slug,omitempty"`
	UserID          int64  `url:"user_id,omitempty"`
	ScreenName      string `url:"screen_name,omitempty"`
	OwnerScreenName string `url:"owner_screen_name,omitempty"`
	OwnerID         int64  `url:"slug,omitempty"`
}

// Create - Add a member to a list. The authenticated user must own the list to be able to add members to it.
// Note that lists cannot have more than 5,000 members.
// https://developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/post-lists-members-create
func (s *ListsService) MembersCreate(params *FavoriteCreateParams) (*List, *http.Response, error) {

	//TODO: The twitter api docs lacks the expected response. Validate that it returns a List or ??
	list := new(List)

	apiError := new(APIError)
	resp, err := s.sling.New().Post("create.json").QueryStruct(params).Receive(list, apiError)
	return list, resp, relevantError(err, *apiError)
}
*/

/*
GET lists/list
GET lists/members
GET lists/members/show

GET lists/memberships
GET lists/ownerships
GET lists/subscriptions
GET lists/show
GET lists/statuses

GET lists/subscribers
GET lists/subscribers/show

*/
