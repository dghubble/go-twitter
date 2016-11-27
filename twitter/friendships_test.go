package twitter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFriendshipService_Create(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/friendships/create.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostForm(t, map[string]string{"user_id": "12345"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"id": 12345, "name": "Doug Williams"}`)
	})

	client := NewClient(httpClient)
	params := &FriendshipCreateParams{UserID: 12345}
	user, _, err := client.Friendships.Create(params)
	assert.Nil(t, err)
	expected := &User{ID: 12345, Name: "Doug Williams"}
	assert.Equal(t, expected, user)
}

func TestFriendshipService_Destroy(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/friendships/destroy.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostForm(t, map[string]string{"user_id": "12345"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"id": 12345, "name": "Doug Williams"}`)
	})

	client := NewClient(httpClient)
	params := &FriendshipDestroyParams{UserID: 12345}
	user, _, err := client.Friendships.Destroy(params)
	assert.Nil(t, err)
	expected := &User{ID: 12345, Name: "Doug Williams"}
	assert.Equal(t, expected, user)
}
