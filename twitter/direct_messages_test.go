package twitter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDMID int64 = 240136858829479936
var testDMIDStr = "240136858829479936"
var testDMJSON = `{
    "id": 240136858829479936,
    "recipient": {
        "screen_name": "s0c1alm3dia"
    },
    "sender": {
        "screen_name": "theSeanCook"
    },
    "text": "booyakasha"
}`
var testDM = DirectMessage{ID: testDMID, Recipient: &User{ScreenName: "s0c1alm3dia"}, Sender: &User{ScreenName: "theSeanCook"}, Text: "booyakasha"}

func TestDirectMessageService_Show(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/direct_messages/show.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"id": testDMIDStr}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[`+testDMJSON+`]`)
	})

	client := NewClient(httpClient)
	dms, _, err := client.DirectMessages.Show(testDMID)
	expected := []DirectMessage{testDM}
	assert.Nil(t, err)
	assert.Equal(t, expected, dms)
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

func TestDirectMessageService_GetNilParams(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/direct_messages.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[`+testDMJSON+`]`)
	})

	client := NewClient(httpClient)
	dms, _, err := client.DirectMessages.Get(nil)
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

func TestDirectMessageService_SentNilParams(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/direct_messages/sent.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[`+testDMJSON+`]`)
	})

	client := NewClient(httpClient)
	dms, _, err := client.DirectMessages.Sent(nil)
	expected := []DirectMessage{testDM}
	assert.Nil(t, err)
	assert.Equal(t, expected, dms)
}

func TestDirectMessageService_SendToID(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/direct_messages/new.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostForm(t, map[string]string{"user_id": "589147592367431680", "text": "booyakasha"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, testDMJSON)
	})

	client := NewClient(httpClient)
	dm, _, err := client.DirectMessages.SendToID(589147592367431680, "booyakasha")
	expected := testDM
	assert.Nil(t, err)
	assert.Equal(t, expected, dm)
}

func TestDirectMessageService_SendToScreenName(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/direct_messages/new.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostForm(t, map[string]string{"screen_name": "s0c1alm3dia", "text": "booyakasha"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, testDMJSON)
	})

	client := NewClient(httpClient)
	dm, _, err := client.DirectMessages.SendToScreenName("s0c1alm3dia", "booyakasha")
	expected := testDM
	assert.Nil(t, err)
	assert.Equal(t, expected, dm)
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
	dm, _, err := client.DirectMessages.Destroy(testDMID, nil)
	expected := testDM
	assert.Nil(t, err)
	assert.Equal(t, expected, dm)
}
