package twitter

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

const (
	streamEndpoint = "tweets/search/"
)

type StreamServiceV2 struct {
	client *http.Client
	sling *sling.Sling
}


func newStreamServiceV2(client *http.Client, sling *sling.Sling) *StreamServiceV2 {
	return &StreamServiceV2{
		client: client,
		sling: sling.Path(streamEndpoint),
	}
}

func (srv *StreamServiceV2) Connect(params *StreamV2FilterParams) (*Stream, error) {
	// req, err := srv.public.New().Get("firehose.json").QueryStruct(params).Request()
	req, err := srv.sling.New().Get("stream").QueryStruct(params).Request()
	fmt.Println(req.URL.Path)
	if err != nil {
		return nil, err
	}
	return newStream(srv.client, req), nil
}

type StreamV2FilterParams struct {
	Expansions  []string `url:"expansions,omitempty,comma"`
	MediaFields []string `url:"media.fields,omitempty,comma"`
	PlaceFields []string `url:"place.fields,omitempty,comma"`
	PollFields  []string `url:"poll.fields,omitempty,comma"`
	TweetFields []string `url:"tweet.fields,omitempty,comma"`
	UserFields  []string `url:"user.fields,omitempty,comma"`
}

type StreamData struct {
	Tweet          *Tweet `json:"data,omitempty"`
	MatchingRules []struct {
		Id  string `json:"id,omitempty"`
		Tag string `json:"tag,omitempty"`

 } `json:"matching_rules,omitempty"`
}
