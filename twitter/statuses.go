package twitter

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// Tweet represents a Twitter Tweet, previously called a status.
// https://dev.twitter.com/overview/api/tweets
// Unused or deprecated fields not provided: Geo, Annotations
// TODO: Place
type Tweet struct {
	Contributors         []Contributor          `json:"contributors"`
	Coordinates          *Coordinates           `json:"coordinates"`
	CreatedAt            string                 `json:"created_at"`
	CurrentUserRetweet   *TweetIdentifier       `json:"current_user_retweet"`
	Entities             *Entities              `json:"entities"`
	FavoriteCount        int                    `json:"favorite_count"`
	Favorited            bool                   `json:"favorited"`
	FilterLevel          string                 `json:"filter_level"`
	ID                   int64                  `json:"id"`
	IDStr                string                 `json:"id_str"`
	InReplyToScreenName  string                 `json:"in_reply_to_screen_name"`
	InReplyToStatusID    int64                  `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr string                 `json:"in_reply_to_status_id_str"`
	InReplyToUserID      int64                  `json:"in_reply_to_user_id"`
	InReplyToUserIDStr   string                 `json:"in_reply_to_user_id_str"`
	Lang                 string                 `json:"lang"`
	PossiblySensitive    bool                   `json:"possibly_sensitive"`
	RetweetCount         int                    `json:"retweet_count"`
	Retweeted            bool                   `json:"retweeted"`
	RetweetedStatus      *Tweet                 `json:"retweeted_status"`
	Source               string                 `json:"source"`
	Scopes               map[string]interface{} `json:"scopes"`
	Text                 string                 `json:"text"`
	Truncated            bool                   `json:"truncated"`
	User                 *User                  `json:"user"`
	WithheldCopyright    bool                   `json:"withheld_copyright"`
	WithheldInCountries  []string               `json:"withheld_in_countries"`
	WithheldScope        string                 `json:"withheld_scope"`
	ExtendedEntities     *ExtendedEntity        `json:"extended_entities"`
	QuotedStatusID       int64                  `json:"quoted_status_id"`
	QuotedStatusIDStr    string                 `json:"quoted_status_id_str"`
	QuotedStatus         *Tweet                 `json:"quoted_status"`
}

// Contributor represents a brief summary of a User identifiers.
type Contributor struct {
	ID         int64  `json:"id"`
	IDStr      string `json:"id_str"`
	ScreenName string `json:"screen_name"`
}

// Coordinates are pairs of longitude and latitude locations.
type Coordinates struct {
	Coordinates [2]float64 `json:"coordinates"`
	Type        string     `json:"type"`
}

// TweetIdentifier represents the id by which a Tweet can be identified.
type TweetIdentifier struct {
	ID    int64  `json:"id"`
	IDStr string `json:"id_str"`
}

// StatusService provides methods for accessing Twitter status API endpoints.
type StatusService struct {
	sling *sling.Sling
}

// newStatusService returns a new StatusService.
func newStatusService(sling *sling.Sling) *StatusService {
	return &StatusService{
		sling: sling.Path("statuses/"),
	}
}

// StatusShowParams are the parameters for StatusService.Show
type StatusShowParams struct {
	ID               int64 `url:"id,omitempty"`
	TrimUser         *bool `url:"trim_user,omitempty"`
	IncludeMyRetweet *bool `url:"include_my_retweet,omitempty"`
	IncludeEntities  *bool `url:"include_entities,omitempty"`
}

// Show returns the requested Tweet.
// https://dev.twitter.com/rest/reference/get/statuses/show/%3Aid
func (s *StatusService) Show(id int64, params *StatusShowParams) (*Tweet, *http.Response, error) {
	if params == nil {
		params = &StatusShowParams{}
	}
	params.ID = id
	tweet := new(Tweet)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("show.json").QueryStruct(params).Receive(tweet, apiError)
	return tweet, resp, relevantError(err, *apiError)
}

// StatusLookupParams are the parameters for StatusService.Lookup
type StatusLookupParams struct {
	ID              []int64 `url:"id,omitempty,comma"`
	TrimUser        *bool   `url:"trim_user,omitempty"`
	IncludeEntities *bool   `url:"include_entities,omitempty"`
	Map             *bool   `url:"map,omitempty"`
}

// Lookup returns the requested Tweets as a slice. Combines ids from the
// required ids argument and from params.Id.
// https://dev.twitter.com/rest/reference/get/statuses/lookup
func (s *StatusService) Lookup(ids []int64, params *StatusLookupParams) ([]Tweet, *http.Response, error) {
	if params == nil {
		params = &StatusLookupParams{}
	}
	params.ID = append(params.ID, ids...)
	tweets := new([]Tweet)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("lookup.json").QueryStruct(params).Receive(tweets, apiError)
	return *tweets, resp, relevantError(err, *apiError)
}

// StatusUpdateParams are the parameters for StatusService.Update
type StatusUpdateParams struct {
	Status             string   `url:"status,omitempty"`
	InReplyToStatusID  int64    `url:"in_reply_to_status_id,omitempty"`
	PossiblySensitive  *bool    `url:"possibly_sensitive,omitempty"`
	Lat                *float64 `url:"lat,omitempty"`
	Long               *float64 `url:"long,omitempty"`
	PlaceID            string   `url:"place_id,omitempty"`
	DisplayCoordinates *bool    `url:"display_coordinates,omitempty"`
	TrimUser           *bool    `url:"trim_user,omitempty"`
	MediaIds           []int64  `url:"media_ids,omitempty,comma"`
}

// Update updates the user's status, also known as Tweeting.
// Requires a user auth context.
// https://dev.twitter.com/rest/reference/post/statuses/update
func (s *StatusService) Update(status string, params *StatusUpdateParams) (*Tweet, *http.Response, error) {
	if params == nil {
		params = &StatusUpdateParams{}
	}
	params.Status = status
	tweet := new(Tweet)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("update.json").BodyForm(params).Receive(tweet, apiError)
	return tweet, resp, relevantError(err, *apiError)
}

// StatusRetweetParams are the parameters for StatusService.Retweet
type StatusRetweetParams struct {
	ID       int64 `url:"id,omitempty"`
	TrimUser *bool `url:"trim_user,omitempty"`
}

// Retweet retweets the Tweet with the given id and returns the original Tweet
// with embedded retweet details.
// Requires a user auth context.
// https://dev.twitter.com/rest/reference/post/statuses/retweet/%3Aid
func (s *StatusService) Retweet(id int64, params *StatusRetweetParams) (*Tweet, *http.Response, error) {
	if params == nil {
		params = &StatusRetweetParams{}
	}
	params.ID = id
	tweet := new(Tweet)
	apiError := new(APIError)
	path := fmt.Sprintf("retweet/%d.json", params.ID)
	resp, err := s.sling.New().Post(path).BodyForm(params).Receive(tweet, apiError)
	return tweet, resp, relevantError(err, *apiError)
}

// StatusDestroyParams are the parameters for StatusService.Destroy
type StatusDestroyParams struct {
	ID       int64 `url:"id,omitempty"`
	TrimUser *bool `url:"trim_user,omitempty"`
}

// Destroy deletes the Tweet with the given id and returns it if successful.
// Requires a user auth context.
// https://dev.twitter.com/rest/reference/post/statuses/destroy/%3Aid
func (s *StatusService) Destroy(id int64, params *StatusDestroyParams) (*Tweet, *http.Response, error) {
	if params == nil {
		params = &StatusDestroyParams{}
	}
	params.ID = id
	tweet := new(Tweet)
	apiError := new(APIError)
	path := fmt.Sprintf("destroy/%d.json", params.ID)
	resp, err := s.sling.New().Post(path).BodyForm(params).Receive(tweet, apiError)
	return tweet, resp, relevantError(err, *apiError)
}

// OEmbedTweet represents a Tweet in oEmbed format.
type OEmbedTweet struct {
	URL          string `json:"url"`
	ProviderURL  string `json:"provider_url"`
	ProviderName string `json:"provider_name"`
	AuthorName   string `json:"author_name"`
	Version      string `json:"version"`
	AuthorURL    string `json:"author_url"`
	Type         string `json:"type"`
	HTML         string `json:"html"`
	Height       int64  `json:"height"`
	Width        int64  `json:"width"`
	CacheAge     string `json:"cache_age"`
}

// StatusOEmbedParams are the parameters for StatusService.OEmbed
type StatusOEmbedParams struct {
	ID         int64  `url:"id,omitempty"`
	URL        string `url:"url,omitempty"`
	Align      string `url:"align,omitempty"`
	MaxWidth   int64  `url:"maxwidth,omitempty"`
	HideMedia  *bool  `url:"hide_media,omitempty"`
	HideThread *bool  `url:"hide_media,omitempty"`
	OmitScript *bool  `url:"hide_media,omitempty"`
	WidgetType string `url:"widget_type,omitempty"`
	HideTweet  *bool  `url:"hide_tweet,omitempty"`
}

// OEmbed returns the requested Tweet in oEmbed format.
// https://dev.twitter.com/rest/reference/get/statuses/oembed
func (s *StatusService) OEmbed(params *StatusOEmbedParams) (*OEmbedTweet, *http.Response, error) {
	oEmbedTweet := new(OEmbedTweet)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("oembed.json").QueryStruct(params).Receive(oEmbedTweet, apiError)
	return oEmbedTweet, resp, relevantError(err, *apiError)
}
