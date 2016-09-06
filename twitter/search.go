package twitter

import (
	"net/http"
	"net/url"

	"github.com/dghubble/sling"
)

// SearchService provides access to Twitter's search
type SearchService struct {
	sling *sling.Sling
}

//
func newSearchService(sling *sling.Sling) *SearchService {
	return &SearchService{
		sling: sling.Path("search/"),
	}
}

// Search is the return results from a search query
type Search struct {
	SearchMetaData SearchMetaData `json:"search_metadata"`
	Statuses       []*Tweet       `json:"statuses"`
}

// SearchMetaData is the metadata related to a search query
type SearchMetaData struct {
	CompletedIn float64 `json:"completed_in"`
	Count       int64   `json:"count"`
	MaxID       int64   `json:"max_id"`
	MaxIDStr    string  `json:"max_id_str"`
	NextResults string  `json:"next_results"`
	Query       string  `json:"query"`
	RefreshURL  string  `json:"refresh_url"`
	SinceID     int64   `json:"since_id"`
	SinceIDStr  string  `json:"since_id_str"`
}

// SearchParams are the parameters for SearchService.Searchs
type SearchParams struct {
	Query           string `url:"q,omitempty"`
	Geocode         string `url:"geocode,omitempty"`
	Lang            string `url:"lang,omitempty"`
	Locale          string `url:"locale,omitempty"`
	ResultType      string `url:"result_type,omitempty"`
	Count           int64  `url:"count,omitempty"`
	Until           string `url:"until,omitempty"`
	SinceID         int64  `url:"since_id,omitempty"`
	MaxID           int64  `url:"max_id,omitempty"`
	IncludeEntities bool   `url:"include_entities,omitempty"`
	NextResults     bool   `url:"q,omitempty"`
}

// Search returns a cursored collection of user ids following the specified user.
// https://dev.twitter.com/rest/reference/get/friends/ids
func (s *SearchService) Search(params *SearchParams) (*Search, *http.Response, error) {
	params.Query = url.QueryEscape(params.Query)
	ids := new(Search)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("tweets.json").QueryStruct(params).Receive(ids, apiError)
	return ids, resp, relevantError(err, *apiError)
}
