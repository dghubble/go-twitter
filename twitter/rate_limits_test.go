package twitter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRateLimitService_Status(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/application/rate_limit_status.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"resources": "statuses,users"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"rate_limit_context":{"access_token":"a_fake_access_token"},"resources":{"statuses":{"/statuses/mentions_timeline":{"limit":75,"remaining":75,"reset":1403602426},"/statuses/lookup":{"limit":900,"remaining":900,"reset":1403602426}}}}`)
	})

	client := NewClient(httpClient)
	rateLimits, _, err := client.RateLimits.Status(&RateLimitParams{Resources: []string{"statuses", "users"}})
	expected := &RateLimit{
		RateLimitContext: &RateLimitContext{AccessToken: "a_fake_access_token"},
		Resources: &RateLimitResources{
			Statuses: map[string]*RateLimitResource{
				"/statuses/mentions_timeline": &RateLimitResource{
					Limit:     75,
					Remaining: 75,
					Reset:     1403602426},
				"/statuses/lookup": &RateLimitResource{
					Limit:     900,
					Remaining: 900,
					Reset:     1403602426}}}}

	assert.Nil(t, err)
	assert.Equal(t, expected, rateLimits)
}
