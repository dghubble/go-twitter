package twitter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFavoriteService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/favorites/list.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"user_id": "113419064", "since_id": "101492475", "include_entities": "false"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{"text": "Gophercon talks!"}, {"text": "Why gophers are so adorable"}]`)
	})

	client := NewClient(httpClient)
	tweets, _, err := client.Favorites.List(&FavoriteListParams{UserID: 113419064, SinceID: 101492475, IncludeEntities: Bool(false)})
	expected := []Tweet{Tweet{Text: "Gophercon talks!"}, Tweet{Text: "Why gophers are so adorable"}}
	assert.Nil(t, err)
	assert.Equal(t, expected, tweets)
}
