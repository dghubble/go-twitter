package twitter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFriendshipService_DestroyHandlesUserIDParam(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/friendships/destroy.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertQuery(t, map[string]string{"user_id": "123"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"id": 123, "screen_name": "croaky"}`)
	})

	client := NewClient(httpClient)
	params := &FriendshipDestroyParams{UserID: 123}
	user, _, err := client.Friendships.Destroy(params)
	expected := &User{ID: 123, ScreenName: "croaky"}
	assert.Nil(t, err)
	assert.Equal(t, expected, user)
}

func TestFriendshipService_DestroyHandlesScreenNameParam(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/friendships/destroy.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertQuery(t, map[string]string{"screen_name": "croaky"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"id": 123, "screen_name": "croaky"}`)
	})

	client := NewClient(httpClient)
	params := &FriendshipDestroyParams{ScreenName: "croaky"}
	tweet, _, err := client.Friendships.Destroy(params)
	expected := &User{ID: 123, ScreenName: "croaky"}
	assert.Nil(t, err)
	assert.Equal(t, expected, tweet)
}

func TestFriendshipService_DestroyHandlesNilParams(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/friendships/destroy.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertQuery(t, map[string]string{}, r)
	})

	client := NewClient(httpClient)
	client.Friendships.Destroy(nil)
}

func TestFriendshipService_APIError(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/friendships/destroy.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		fmt.Fprintf(w, `{"errors": [{"message": "Sorry, that page does not exist.", "code": 34}]}`)
	})

	client := NewClient(httpClient)
	_, _, err := client.Friendships.Destroy(nil)
	expected := APIError{
		Errors: []ErrorDetail{
			ErrorDetail{Message: "Sorry, that page does not exist.", Code: 34},
		},
	}
	if assert.Error(t, err) {
		assert.Equal(t, expected, err)
	}
}
