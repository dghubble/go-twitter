package twitter

import (
	"github.com/dghubble/sling"
	"net/http"
)

// DirectMessage is a direct message to a single recipient.
type DirectMessage struct {
	CreatedAt           string    `json:"created_at"`
	Entities            *Entities `json:"entities"`
	ID                  int64     `json:"id"`
	IDStr               string    `json:"id_str"`
	Recipient           *User     `json:"recipient"`
	RecipientID         int64     `json:"recipient_id"`
	RecipientScreenName string    `json:"recipient_screen_name"`
	Sender              *User     `json:"sender"`
	SenderID            int64     `json:"sender_id"`
	SenderScreenName    string    `json:"sender_screen_name"`
	Text                string    `json:"text"`
}

// DirectMessageService provides methods for accessing Twitter status API endpoints.
type DirectMessageService struct {
	baseSling *sling.Sling
	sling     *sling.Sling
}

// newDirectMessageService returns a new StatusService.
func newDirectMessageService(sling *sling.Sling) *DirectMessageService {
	return &DirectMessageService{
		baseSling: sling.New(),
		sling:     sling.Path("direct_messages/"),
	}
}

// DirectMessageShowParams are the parameters for DirectMessageService.Show
type DirectMessageShowParams struct {
	ID int64 `url:"id,omitempty"`
}

// Show returns the requested DirectMessage and all response messages.
// https://dev.twitter.com/rest/reference/get/direct_messages/show
func (s *DirectMessageService) Show(id int64) ([]DirectMessage, *http.Response, error) {
	params := &StatusShowParams{ID: id}
	dms := new([]DirectMessage)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("show.json").QueryStruct(params).Receive(dms, apiError)
	return *dms, resp, relevantError(err, *apiError)
}

// DirectMessageGetParams are the parameters for DirectMessageService.Get
type DirectMessageGetParams struct {
	SinceID         int64 `url:"since_id,omitempty"`
	MaxID           int64 `url:"max_id,omitempty"`
	Count           int   `url:"count,omitempty"`
	IncludeEntities *bool `url:"include_entities,omitempty"`
	SkipStatus      *bool `url:"skip_status,omitempty"`
}

// Get returns the all the direct messages
// https://dev.twitter.com/rest/reference/get/direct_messages
func (s *DirectMessageService) Get(params *DirectMessageGetParams) ([]DirectMessage, *http.Response, error) {
	if params == nil {
		params = &DirectMessageGetParams{}
	}
	dms := new([]DirectMessage)
	apiError := new(APIError)
	resp, err := s.baseSling.New().Get("direct_messages.json").QueryStruct(params).Receive(dms, apiError)
	return *dms, resp, relevantError(err, *apiError)
}

// DirectMessageSentParams are the parameters for DirectMessageService.Sent
type DirectMessageSentParams struct {
	SinceID         int64 `url:"since_id,omitempty"`
	MaxID           int64 `url:"max_id,omitempty"`
	Count           int   `url:"count,omitempty"`
	Page            int   `url:"page,omitempty"`
	IncludeEntities *bool `url:"include_entities,omitempty"`
}

// Sent returns the all the direct messages sent
// https://dev.twitter.com/rest/reference/get/direct_messages/sent
func (s *DirectMessageService) Sent(params *DirectMessageSentParams) ([]DirectMessage, *http.Response, error) {
	if params == nil {
		params = &DirectMessageSentParams{}
	}
	dms := new([]DirectMessage)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("sent.json").QueryStruct(params).Receive(dms, apiError)
	return *dms, resp, relevantError(err, *apiError)
}

// DirectMessageNewParams are the parameters for DirectMessageService.New
type DirectMessageNewParams struct {
	UserID     int64  `url:"user_id,omitempty"`
	ScreenName string `url:"screen_name,omitempty"`
	Text       string `url:"text"`
}

// New creates a new direct message
// https://dev.twitter.com/rest/reference/get/direct_messages/new
func (s *DirectMessageService) New(params DirectMessageNewParams) (DirectMessage, *http.Response, error) {
	dm := new(DirectMessage)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("new.json").BodyForm(params).Receive(dm, apiError)
	return *dm, resp, relevantError(err, *apiError)
}

// SendToID sends a direct message by twitter user id
func (s *DirectMessageService) SendToID(userID int64, text string) (DirectMessage, *http.Response, error) {
	return s.New(DirectMessageNewParams{UserID: userID, Text: text})
}

// SendToScreenName sends a message by twitter user id
func (s *DirectMessageService) SendToScreenName(screenName, text string) (DirectMessage, *http.Response, error) {
	return s.New(DirectMessageNewParams{ScreenName: screenName, Text: text})
}

// DirectMessageDestroyParams are the parameters for DirectMessageService.Destroy
type DirectMessageDestroyParams struct {
	ID              int64 `url:"id,omitempty"`
	IncludeEntities *bool `url:"include_entities,omitempty"`
}

// Destroy deletes a direct message
// https://dev.twitter.com/rest/reference/get/direct_messages/new
func (s *DirectMessageService) Destroy(id int64, params *DirectMessageDestroyParams) (DirectMessage, *http.Response, error) {
	if params == nil {
		params = &DirectMessageDestroyParams{}
	}
	params.ID = id
	dm := new(DirectMessage)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("destroy.json").BodyForm(params).Receive(dm, apiError)
	return *dm, resp, relevantError(err, *apiError)
}
