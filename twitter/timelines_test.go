package twitter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimelineService_UserTimeline(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/statuses/user_timeline.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"user_id": "113419064", "trim_user": "true", "include_rts": "false"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"text": "Gophercon talks!"}, {"text": "Why gophers are so adorable"}]`)
	})

	client := NewClient(httpClient)
	tweets, _, err := client.Timelines.UserTimeline(&UserTimelineParams{UserID: 113419064, TrimUser: Bool(true), IncludeRetweets: Bool(false)})
	expected := []Tweet{{Text: "Gophercon talks!"}, {Text: "Why gophers are so adorable"}}
	assert.Nil(t, err)
	assert.Equal(t, expected, tweets)
}

func TestTimelineService_HomeTimeline(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/statuses/home_timeline.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"since_id": "589147592367431680", "exclude_replies": "false"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"text": "Live on #Periscope"}, {"text": "Clickbait journalism"}, {"text": "Useful announcement"}]`)
	})

	client := NewClient(httpClient)
	tweets, _, err := client.Timelines.HomeTimeline(&HomeTimelineParams{SinceID: 589147592367431680, ExcludeReplies: Bool(false)})
	expected := []Tweet{{Text: "Live on #Periscope"}, {Text: "Clickbait journalism"}, {Text: "Useful announcement"}}
	assert.Nil(t, err)
	assert.Equal(t, expected, tweets)
}

func TestTimelineService_MentionTimeline(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/statuses/mentions_timeline.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"count": "20", "include_entities": "false"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"text": "@dghubble can I get verified?"}, {"text": "@dghubble why are gophers so great?"}]`)
	})

	client := NewClient(httpClient)
	tweets, _, err := client.Timelines.MentionTimeline(&MentionTimelineParams{Count: 20, IncludeEntities: Bool(false)})
	expected := []Tweet{{Text: "@dghubble can I get verified?"}, {Text: "@dghubble why are gophers so great?"}}
	assert.Nil(t, err)
	assert.Equal(t, expected, tweets)
}

func TestTimelineService_RetweetsOfMeTimeline(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/statuses/retweets_of_me.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"trim_user": "false", "include_user_entities": "false"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"text": "RT Twitter UK edition"}, {"text": "RT Triply-replicated Gophers"}]`)
	})

	client := NewClient(httpClient)
	tweets, _, err := client.Timelines.RetweetsOfMeTimeline(&RetweetsOfMeTimelineParams{TrimUser: Bool(false), IncludeUserEntities: Bool(false)})
	expected := []Tweet{{Text: "RT Twitter UK edition"}, {Text: "RT Triply-replicated Gophers"}}
	assert.Nil(t, err)
	assert.Equal(t, expected, tweets)
}
