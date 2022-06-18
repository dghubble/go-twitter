package twitter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testDMEvent = DirectMessageEvent{
		CreatedAt: "1542410751275",
		ID:        "1063573894173323269",
		Type:      "message_create",
		Message: &DirectMessageEventMessage{
			SenderID: "623265148",
			Target: &DirectMessageTarget{
				RecipientID: "3694959333",
			},
			Data: &DirectMessageData{
				Text: "example",
				Entities: &Entities{
					Hashtags:     []HashtagEntity{},
					Urls:         []URLEntity{},
					UserMentions: []MentionEntity{},
					Symbols:      []SymbolEntity{},
				},
			},
		},
	}
	testDMEventID   = "1063573894173323269"
	testDMEventJSON = `
{
	"type": "message_create",
	"id": "1063573894173323269",
	"created_timestamp": "1542410751275",
	"message_create": {
		"target": {
			"recipient_id": "3694959333"
		},
		"sender_id": "623265148",
		"message_data": {
			"text": "example",
			"entities": {
				"hashtags": [],
				"symbols": [],
				"user_mentions": [],
				"urls": []
			}
		}
  }
}`
	testDMEventShowJSON     = `{"event": ` + testDMEventJSON + `}`
	testDMEventListJSON     = `{"events": [` + testDMEventJSON + `], "next_cursor": "AB345dkfC"}`
	testDMEventNewInputJSON = `{"event":{"type":"message_create","message_create":{"target":{"recipient_id":"3694959333"},"message_data":{"text":"example","entities":{"hashtags":null,"media":null,"urls":null,"user_mentions":null,"symbols":null,"polls":null}}}}}
`

	// DEPRECATED
	testDM = DirectMessage{
		ID:        240136858829479936,
		Recipient: &User{ScreenName: "theSeanCook"},
		Sender:    &User{ScreenName: "s0c1alm3dia"},
		Text:      "hello world",
	}
	testDMIDStr = "240136858829479936"
	testDMJSON  = `{"id": 240136858829479936,"recipient": {"screen_name": "theSeanCook"},"sender": {"screen_name": "s0c1alm3dia"},"text": "hello world"}`
)

func TestDirectMessageService_EventsNew(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/direct_messages/events/new.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostJSON(t, testDMEventNewInputJSON, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, testDMEventShowJSON)
	})

	client := NewClient(httpClient)
	event, _, err := client.DirectMessages.EventsNew(&DirectMessageEventsNewParams{
		Event: &DirectMessageEvent{
			Type: "message_create",
			Message: &DirectMessageEventMessage{
				Target: &DirectMessageTarget{
					RecipientID: "3694959333",
				},
				Data: &DirectMessageData{
					Text:     "example",
					Entities: &Entities{},
				},
			},
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, &testDMEvent, event)
}

func TestDirectMessageService_EventsShow(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/direct_messages/events/show.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"id": testDMEventID}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, testDMEventShowJSON)
	})

	client := NewClient(httpClient)
	event, _, err := client.DirectMessages.EventsShow(testDMEventID, nil)
	assert.Nil(t, err)
	assert.Equal(t, &testDMEvent, event)
}

func TestDirectMessageService_EventsList(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/direct_messages/events/list.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"count": "10"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, testDMEventListJSON)
	})
	expected := &DirectMessageEvents{
		Events:     []DirectMessageEvent{testDMEvent},
		NextCursor: "AB345dkfC",
	}

	client := NewClient(httpClient)
	events, _, err := client.DirectMessages.EventsList(&DirectMessageEventsListParams{Count: 10})
	assert.Equal(t, expected, events)
	assert.Nil(t, err)
}

func TestDirectMessageService_EventsDestroy(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/direct_messages/events/destroy.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "DELETE", r)
		assertQuery(t, map[string]string{"id": testDMEventID}, r)
		w.Header().Set("Content-Type", "application/json")
		// successful delete returns 204 No Content
		w.WriteHeader(204)
	})

	client := NewClient(httpClient)
	resp, err := client.DirectMessages.EventsDestroy(testDMEventID)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestDirectMessageService_EventsDestroyError(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/direct_messages/events/destroy.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "DELETE", r)
		assertQuery(t, map[string]string{"id": testDMEventID}, r)
		w.Header().Set("Content-Type", "application/json")
		// failure to delete event that doesn't exist
		w.WriteHeader(404)
		fmt.Fprintf(w, `{"errors":[{"code": 34, "message": "Sorry, that page does not exist"}]}`)
	})
	expected := APIError{
		Errors: []ErrorDetail{
			ErrorDetail{Code: 34, Message: "Sorry, that page does not exist"},
		},
	}

	client := NewClient(httpClient)
	resp, err := client.DirectMessages.EventsDestroy(testDMEventID)
	assert.NotNil(t, resp)
	if assert.Error(t, err) {
		assert.Equal(t, expected, err)
	}
}

// DEPRECATED

func TestDirectMessageService_Show(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/direct_messages/show.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"id": testDMIDStr}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, testDMJSON)
	})

	client := NewClient(httpClient)
	dms, _, err := client.DirectMessages.Show(testDM.ID)
	assert.Nil(t, err)
	assert.Equal(t, &testDM, dms)
}

func TestDirectMessageService_Get(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/direct_messages.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"since_id": "589147592367431680", "count": "1"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[`+testDMJSON+`]`)
	})

	client := NewClient(httpClient)
	params := &DirectMessageGetParams{SinceID: 589147592367431680, Count: 1}
	dms, _, err := client.DirectMessages.Get(params)
	expected := []DirectMessage{testDM}
	assert.Nil(t, err)
	assert.Equal(t, expected, dms)
}

func TestDirectMessageService_Sent(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/direct_messages/sent.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"since_id": "589147592367431680", "count": "1"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[`+testDMJSON+`]`)
	})

	client := NewClient(httpClient)
	params := &DirectMessageSentParams{SinceID: 589147592367431680, Count: 1}
	dms, _, err := client.DirectMessages.Sent(params)
	expected := []DirectMessage{testDM}
	assert.Nil(t, err)
	assert.Equal(t, expected, dms)
}

func TestDirectMessageService_New(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/direct_messages/new.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostForm(t, map[string]string{"screen_name": "theseancook", "text": "hello world"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, testDMJSON)
	})

	client := NewClient(httpClient)
	params := &DirectMessageNewParams{ScreenName: "theseancook", Text: "hello world"}
	dm, _, err := client.DirectMessages.New(params)
	assert.Nil(t, err)
	assert.Equal(t, &testDM, dm)
}

func TestDirectMessageService_Destroy(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/direct_messages/destroy.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostForm(t, map[string]string{"id": testDMIDStr}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, testDMJSON)
	})

	client := NewClient(httpClient)
	dm, _, err := client.DirectMessages.Destroy(testDM.ID, nil)
	assert.Nil(t, err)
	assert.Equal(t, &testDM, dm)
}
