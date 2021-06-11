package twitter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestBlockService_CreateService(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/blocks/create.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertQuery(t, map[string]string{"screen_name": "golang"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"screen_name": "golang"}`)
	})

	client := NewClient(httpClient)
	users, _, err := client.Block.Create(&BlockCreateParams{ScreenName: "golang"})
	expected := User{ScreenName: "golang"}
	assert.Nil(t, err)
	assert.Equal(t, expected, users)
}

func TestBlockService_DestroyService(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/blocks/destroy.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertQuery(t, map[string]string{"screen_name": "golang"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"screen_name": "golang"}`)
	})

	client := NewClient(httpClient)
	users, _, err := client.Block.Destroy(&BlockDestroyParams{ScreenName: "golang"})
	expected := User{ScreenName: "golang"}
	assert.Nil(t, err)
	assert.Equal(t, expected, users)
}
