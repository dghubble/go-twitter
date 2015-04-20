package twitter

import (
	"github.com/dghubble/sling"
	"net/http"
)

// https://dev.twitter.com/overview/api/tweets
// Unused or deprecated fields not provided: Geo, Annotations
// TODO: Place
type Tweet struct {
	Contributors         []Contributor     `json:"contributors"`
	Coordinates          *Coordinates      `json:"coordinates"`
	CreatedAt            string            `json:"created_at"`
	CurrentUserRetweet   *TweetIdentifier  `json:"current_user_retweet"`
	Entities             *Entities         `json:"entities"`
	FavoriteCount        int               `json:"favorite_count"`
	Favorited            bool              `json:"favorited"`
	FilterLevel          string            `json:"filter_level"`
	Id                   int64             `json:"id"`
	IdStr                string            `json:"id_str"`
	InReplyToScreenName  string            `json:"in_reply_to_screen_name"`
	InReplyToStatusId    int64             `json:"in_reply_to_status_id"`
	InReplyToStatusIdStr string            `json:"in_reply_to_status_id_str"`
	InReplyToUserId      int64             `json:"in_reply_to_user_id"`
	InReplyToUserIdStr   string            `json:"in_reply_to_user_id_str"`
	Lang                 string            `json:"lang"`
	PossiblySensitive    bool              `json:"possibly_sensitive"`
	RetweetCount         int               `json:"retweet_count"`
	Retweeted            bool              `json:"retweeted"`
	RetweetedStatus      *Tweet            `json:"retweeted_status"`
	Source               string            `json:"source"`
	Scopes               map[string]string `json:"scopes"`
	Text                 string            `json:"text"`
	Truncated            bool              `json:"truncated"`
	User                 *User             `json:"user"`
	WithheldCopyright    bool              `json:"withheld_copyright"`
	WithheldInCountries  []string          `json:"withheld_in_countries"`
	WithheldScope        string            `json:"withheld_scope"`
}

type Contributor struct {
	Id         int64  `json:"id"`
	IdStr      string `json:"id_str"`
	ScreenName string `json:"screen_name"`
}

type Coordinates struct {
	Coordinates [2]float64 `json:"coordinates"`
	Type        string     `json:"type"`
}

type TweetIdentifier struct {
	Id    int64  `json:"id"`
	IdStr string `json:"id_str"`
}

type StatusService struct {
	sling *sling.Sling
}

func NewStatusService(sling *sling.Sling) *StatusService {
	return &StatusService{
		sling: sling.Path("statuses/"),
	}
}

type StatusShowParams struct {
	Id               int64 `url:"id,omitempty"`
	TrimUser         *bool `url:"trim_user,omitempty"`
	IncludeMyRetweet *bool `url:"include_my_retweet,omitempty"`
	IncludeEntities  *bool `url:"include_entities,omitempty"`
}

// Show returns the specified Tweet.
// https://dev.twitter.com/rest/reference/get/statuses/show/%3Aid
func (s *StatusService) Show(id int64, params *StatusShowParams) (*Tweet, *http.Response, error) {
	params.Id = id
	tweet := new(Tweet)
	resp, err := s.sling.New().Get("show.json").QueryStruct(params).Receive(tweet)
	return tweet, resp, err
}

type StatusLookupParams struct {
	Id              []int64 `json:"id,omitempty"`
	TrimUser        *bool   `json:"trim_user,omitempty"`
	IncludeEntities *bool   `json:"include_entities,omitempty"`
	Map             *bool   `json:"map,omitempty"`
}

// Lookup returns the specified slice of Tweets
// https://dev.twitter.com/rest/reference/get/statuses/lookup
func (s *StatusService) Lookup(params *StatusLookupParams) ([]Tweet, *http.Response, error) {
	tweets := new([]Tweet)
	resp, err := s.sling.New().Get("lookup.json").QueryStruct(params).Receive(tweets)
	return *tweets, resp, err
}
