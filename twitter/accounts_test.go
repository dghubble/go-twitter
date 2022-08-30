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

func TestAccountService_UpdateProfileImage(t *testing.T) {
	base64PNG := "iVBORw0KGgoAAAANSUhEUgAAAGQAAABkCAMAAABHPGVmAAABCFBMVEX////+/v7++/z9+Pr99/r89vn89Pj78vf66/P66vL56fH56PD45u/45e/45O734ez23ur23On12uj12ef11+b11uX00+Pz0eLz0OHyzN/wxNrww9nwwtnvwtjuvdbuvNXtutTtudPtuNLtt9LstdHstNDss9Drss/rr83qrMvpqMjoo8XmmsDlmb/lmL7klLzjkbrjj7nii7fiirbhirXghbPghbLghLLfgLDff6/efa7dd6rddqncdqnccqfbcabbb6Xaa6PZaKHYZJ7YYp3XX5vVV5bUVZXTUJHSTZDRSI3QQ4rPP4fPPofOOoTNOYPNN4LMNYHLMX7LMH7KLXzKLHvKK3vKKnp2VKJ7AAAA9ElEQVR42u3JRWIEMQxE0QozMzM2hJmZaajuf5PYHfAcQD2regtb0oeIiIiIiIiIiOSpbTjTh47s7wfQOzffXxUMHDBzjpPsv0Hzpf/jEAwsxAd8i+MxPHE3TdNxnPJsYuZ4KgQTS7xwb2O5UO+3uq9yS1Wwss3EvSN8WVhYGAWeeTcUgpVbTrp3nd4WsFwkDxtCsPHBVvce8SJJkkEA7TtF7oVgopOv/nvg6P9pms8hmJjmtXvri5UmeCt1wCLvf4OVDW66d4ClB6e7i4/7Z59c+w1Wrjjr3lV6xfqeuwr5HoWQj6bejjrUhoiIiIiIiIiIfANKKjEabNmIvgAAAABJRU5ErkJggg=="
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/account/update_profile_image.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostForm(t, map[string]string{"image": base64PNG}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"image": "%s"}`, base64PNG)
	})

	client := NewClient(httpClient)
	params := &AccountUpdateProfileImageParams{Image: base64PNG}
	user, _, err := client.Accounts.UpdateProfileImage(params)
	expected := &User{DefaultProfileImage: false}
	assert.Nil(t, err)
	assert.Equal(t, expected, user)

}
