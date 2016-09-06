package twitter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchService_Search(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/search/tweets.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"q": "sports", "count": "1"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"search_metadata":{"completed_in":0.033,"count":1,"max_id":773094573430104065,"max_id_str":"773094573430104065","next_results":"?max_id=773094573430104064&q=sports&count=1&include_entities=1","query":"sports","refresh_url":"?since_id=773094573430104065&q=sports&include_entities=1","since_id":0,"since_id_str":"0"},"statuses":[{"id":773094573430104065}]}`)
	})
	expected := &Search{
		SearchMetaData: SearchMetaData{
			CompletedIn: 0.033,
			Count:       1,
			MaxID:       773094573430104065,
			MaxIDStr:    "773094573430104065",
			NextResults: "?max_id=773094573430104064&q=sports&count=1&include_entities=1",
			Query:       "sports",
			RefreshURL:  "?since_id=773094573430104065&q=sports&include_entities=1",
			SinceID:     0,
			SinceIDStr:  "0",
		},
		Statuses: []*Tweet{
			&Tweet{
				ID: 773094573430104065,
			},
		},
	}

	client := NewClient(httpClient)
	params := &SearchParams{
		Query: "sports",
		Count: 1,
	}
	searchResult, _, err := client.Search.Search(params)
	assert.Nil(t, err)
	assert.Equal(t, expected, searchResult)
}
