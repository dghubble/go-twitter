package twitter

import (
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

// DirectMessageEvent is a list direct message event
type DirectMessageEvent struct {
	NextCursor string         `json:"next_cursor"`
	Events     []MessageEvent `json:"events"`
	Apps       string         `json:"apps"`
}

// MessageEvent is a signle direct message sent or received
type MessageEvent struct {
	Type      string   `json:"type"`
	ID        int64    `json:"id"`
	CreatedAt string   `json:"created_timestamp"`
	Message   *Message `json:"message_create"`
}

// Message contains the Sender data as well as the Message contents
type Message struct {
	SenderID string       `json:"sender_id"`
	Target   *Target      `json:"target"`
	Data     *MessageData `json:"message_data"`
}

// Target recipient information
type Target struct {
	RecipientID string `json:"recipient_id"`
}

// MessageData contains the raw text of the message sent or received
type MessageData struct {
	Text     string    `json:"text"`
	Entities *Entities `json:"entitites"`
}

// DirectMessageEventsGetParams are the parameters for DirectMessageEvents.Get
type DirectMessageEventsGetParams struct {
	NextCursor string `url:"cursor,omitempty"`
	Count      int    `url:"count,omitempty"`
}

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

// CreatedAtTime is a convenience wrapper that returns the Created_at time, parsed as a time.Time struct
func (d DirectMessage) CreatedAtTime() (time.Time, error) {
	return time.Parse(time.RubyDate, d.CreatedAt)
}

// DirectMessageService provides methods for accessing Twitter direct message
// API endpoints.
type DirectMessageService struct {
	baseSling *sling.Sling
	sling     *sling.Sling
}

// newDirectMessageService returns a new DirectMessageService.
func newDirectMessageService(sling *sling.Sling) *DirectMessageService {
	return &DirectMessageService{
		baseSling: sling.New(),
		sling:     sling.Path("direct_messages/"),
	}
}

// directMessageShowParams are the parameters for DirectMessageService.Show
type directMessageShowParams struct {
	ID int64 `url:"id,omitempty"`
}

// Show returns the requested Direct Message.
// Requires a user auth context with DM scope.
// https://dev.twitter.com/rest/reference/get/direct_messages/show
func (s *DirectMessageService) Show(id int64) (*DirectMessage, *http.Response, error) {
	params := &directMessageShowParams{ID: id}
	dm := new(DirectMessage)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("show.json").QueryStruct(params).Receive(dm, apiError)
	return dm, resp, relevantError(err, *apiError)
}

// DirectMessageGetParams are the parameters for DirectMessageService.Get
type DirectMessageGetParams struct {
	SinceID         int64 `url:"since_id,omitempty"`
	MaxID           int64 `url:"max_id,omitempty"`
	Count           int   `url:"count,omitempty"`
	IncludeEntities *bool `url:"include_entities,omitempty"`
	SkipStatus      *bool `url:"skip_status,omitempty"`
}

// Get returns recent Direct Messages received by the authenticated user.
// Requires a user auth context with DM scope.
// https://dev.twitter.com/rest/reference/get/direct_messages
func (s *DirectMessageService) Get(params *DirectMessageGetParams) ([]DirectMessage, *http.Response, error) {
	dms := new([]DirectMessage)
	apiError := new(APIError)
	resp, err := s.baseSling.New().Get("direct_messages.json").QueryStruct(params).Receive(dms, apiError)
	return *dms, resp, relevantError(err, *apiError)
}

// GetEvents returns recent Direct Message Events received or sent by the authenticated user.
// Requires a user auth context with DM scope.
// https://developer.twitter.com/en/docs/direct-messages/sending-and-receiving/api-reference/list-events
func (s *DirectMessageService) GetEvents(params *DirectMessageEventsGetParams) (DirectMessageEvent, *http.Response, error) {
	event := new(DirectMessageEvent)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("events/list.json").QueryStruct(params).Receive(event, apiError)
	return *event, resp, relevantError(err, *apiError)
}

// DirectMessageSentParams are the parameters for DirectMessageService.Sent
type DirectMessageSentParams struct {
	SinceID         int64 `url:"since_id,omitempty"`
	MaxID           int64 `url:"max_id,omitempty"`
	Count           int   `url:"count,omitempty"`
	Page            int   `url:"page,omitempty"`
	IncludeEntities *bool `url:"include_entities,omitempty"`
}

// Sent returns recent Direct Messages sent by the authenticated user.
// Requires a user auth context with DM scope.
// https://dev.twitter.com/rest/reference/get/direct_messages/sent
func (s *DirectMessageService) Sent(params *DirectMessageSentParams) ([]DirectMessage, *http.Response, error) {
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

// New sends a new Direct Message to a specified user as the authenticated
// user.
// Requires a user auth context with DM scope.
// https://dev.twitter.com/rest/reference/post/direct_messages/new
func (s *DirectMessageService) New(params *DirectMessageNewParams) (*DirectMessage, *http.Response, error) {
	dm := new(DirectMessage)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("new.json").BodyForm(params).Receive(dm, apiError)
	return dm, resp, relevantError(err, *apiError)
}

// DirectMessageDestroyParams are the parameters for DirectMessageService.Destroy
type DirectMessageDestroyParams struct {
	ID              int64 `url:"id,omitempty"`
	IncludeEntities *bool `url:"include_entities,omitempty"`
}

// Destroy deletes the Direct Message with the given id and returns it if
// successful.
// Requires a user auth context with DM scope.
// https://dev.twitter.com/rest/reference/post/direct_messages/destroy
func (s *DirectMessageService) Destroy(id int64, params *DirectMessageDestroyParams) (*DirectMessage, *http.Response, error) {
	if params == nil {
		params = &DirectMessageDestroyParams{}
	}
	params.ID = id
	dm := new(DirectMessage)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("destroy.json").BodyForm(params).Receive(dm, apiError)
	return dm, resp, relevantError(err, *apiError)
}
