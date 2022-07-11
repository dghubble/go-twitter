package twitter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountService_VerifyCredentials(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/account/verify_credentials.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"include_entities": "false", "include_email": "true"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"name": "Dalton Hubble", "id": 623265148}`)
	})

	client := NewClient(httpClient)
	user, _, err := client.Accounts.VerifyCredentials(&AccountVerifyParams{IncludeEntities: Bool(false), IncludeEmail: Bool(true)})
	expected := &User{Name: "Dalton Hubble", ID: 623265148}
	assert.Nil(t, err)
	assert.Equal(t, expected, user)
}
func TestAccountService_UpdateProfile(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/account/update_profile.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertQuery(t, map[string]string{"location": "anywhere"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"name": "xkcdComic", "location":"anywhere"}`)
	})

	client := NewClient(httpClient)
	params := &AccountUpdateProfileParams{Location: "anywhere"}
	user, _, err := client.Accounts.UpdateProfile(params)
	expected := &User{Name: "xkcdComic", Location: "anywhere"}
	assert.Nil(t, err)
	assert.Equal(t, expected, user)
}
