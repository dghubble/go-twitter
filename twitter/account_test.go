package twitter

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAccountService_VerifyCredentials(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/account/verify_credentials.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"include_entities": "false", "include_email": "true"}, r)
		fmt.Fprintf(w, `{"name": "Dalton Hubble", "id": 623265148}`)
	})

	client := NewClient(httpClient)
	user, _, err := client.Accounts.VerifyCredentials(&AccountVerifyParams{IncludeEntities: Bool(false), IncludeEmail: Bool(true)})
	if err != nil {
		t.Errorf("Accounts.VerifyCredentials error %+v", err)
	}
	expected := &User{Name: "Dalton Hubble", ID: 623265148}
	if !reflect.DeepEqual(expected, user) {
		t.Errorf("Accounts.VerifyCredentials expected:\n%+v, got:\n %+v", expected, user)
	}
}
