package twitter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const rateLimitStatusJSON = `{ 
	"rate_limit_context": { "access_token": "foo-bar" },
	"resources": {
		"users": {
			"/users/profile_banner": {
        		"limit": 180,
        		"remaining": 170,
        		"reset": 1403602426
      		}
		},
		"search": {
			"/search/tweets": {
        		"limit": 160,
        		"remaining": 150,
        		"reset": 1403602427
      		}
		}
	}
}`

func TestApplicationService_RateLimitStatus(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/application/rate_limit_status.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"resources": "help,users,search,statuses"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, rateLimitStatusJSON)
	})

	client := NewClient(httpClient)
	user, _, err := client.Application.RateLimitStatus()

	expected := &RateLimitStatus{
		RateLimitContext: &RateLimitContext{
			AccessToken: "foo-bar",
		},
		Resources: &RateLimitResources{
			Users: &UsersRateLimitResource{
				ProfileBanner: &RateLimitResource{
					Limit:     180,
					Remaining: 170,
					Reset:     int64(1403602426),
				},
			},
			Search: &SearchRateLimitResource{
				Tweets: &RateLimitResource{
					Limit:     160,
					Remaining: 150,
					Reset:     int64(1403602427),
				},
			},
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, user)
}
