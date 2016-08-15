package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

// FavoriteService provides methods for accessing Twitter favorite API endpoints.
//
// Note: the like action was known as favorite before November 3, 2015; the
// historical naming remains in API methods and object properties.
type FavoriteService struct {
	sling *sling.Sling
}

// newFavoriteService returns a new FavoriteService.
func newFavoriteService(sling *sling.Sling) *FavoriteService {
	return &FavoriteService{
		sling: sling.Path("favorites/"),
	}
}

// FavoriteListParams are the parameters for FavoriteService.List.
type FavoriteListParams struct {
	UserID          int64  `url:"user_id,omitempty"`
	ScreenName      string `url:"screen_name,omitempty"`
	Count           int    `url:"count,omitempty"`
	SinceID         int64  `url:"since_id,omitempty"`
	MaxID           int64  `url:"max_id,omitempty"`
	IncludeEntities *bool  `url:"include_entities,omitempty"`
}

// List returns liked Tweets from the specified user.
// https://dev.twitter.com/rest/reference/get/favorites/list
func (s *FavoriteService) List(params *FavoriteListParams) ([]Tweet, *http.Response, error) {
	favorites := new([]Tweet)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("list.json").QueryStruct(params).Receive(favorites, apiError)
	return *favorites, resp, relevantError(err, *apiError)
}
