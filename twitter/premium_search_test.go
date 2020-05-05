package twitter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPremiumSearchService_Tweets(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	assertSearchBody := func(t *testing.T, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"query": "url:\"http://example.com\"", "tag": "8HYG54ZGTU", "fromDate": "201512220000", "toDate": "201712220000", "maxResults": "500", "next": "NTcxODIyMDMyODMwMjU1MTA0"}, r)
	}
	setResponse := func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"results":[{"id":781760642139250689}],"next":"NTcxODIyMDMyODMwMjU1MTA0","requestParameters":{"maxResults":500,"fromDate":"201512200000","toDate":"201712200000"}}`)
	}

	mux.HandleFunc("/1.1/tweets/search/fullarchive/test.json", func(w http.ResponseWriter, r *http.Request) {
		assertSearchBody(t, r)
		setResponse(w)
	})
	mux.HandleFunc("/1.1/tweets/search/30day/test.json", func(w http.ResponseWriter, r *http.Request) {
		assertSearchBody(t, r)
		setResponse(w)
	})

	params := &PremiumSearchTweetParams{
		Query:      "url:\"http://example.com\"",
		Tag:        "8HYG54ZGTU",
		FromDate:   "201512220000",
		ToDate:     "201712220000",
		MaxResults: 500,
		Next:       "NTcxODIyMDMyODMwMjU1MTA0",
	}
	expected := &PremiumSearch{
		Results: []Tweet{
			Tweet{ID: 781760642139250689},
		},
		Next: "NTcxODIyMDMyODMwMjU1MTA0",
		RequestParameters: &RequestParameters{
			MaxResults: 500,
			FromDate:   "201512200000",
			ToDate:     "201712200000",
		},
	}

	{
		client := NewClient(httpClient)
		search, _, err := client.PremiumSearch.SearchFullArchive(
			params,
			"test",
		)
		assert.Nil(t, err)
		assert.Equal(t, expected, search)
	}
	{
		client := NewClient(httpClient)
		search, _, err := client.PremiumSearch.Search30Days(
			params,
			"test",
		)
		assert.Nil(t, err)
		assert.Equal(t, expected, search)
	}
}

func TestPremiumSearchService_Counts(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	assertCountBody := func(t *testing.T, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"query": "url:\"http://example.com\"", "tag": "8HYG54ZGTU", "fromDate": "201512220000", "toDate": "201712220000", "bucket": "day", "next": "NTcxODIyMDMyODMwMjU1MTA0"}, r)
	}
	setResponse := func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"results":[{"timePeriod":"201701010000","count":32},{"timePeriod":"201701020000","count":45}],"totalCount":2027,"requestParameters":{"bucket":"day","fromDate":"201512200000","toDate":"201712200000"}}`)
	}

	mux.HandleFunc("/1.1/tweets/search/fullarchive/test/counts.json", func(w http.ResponseWriter, r *http.Request) {
		assertCountBody(t, r)
		setResponse(w)
	})
	mux.HandleFunc("/1.1/tweets/search/30day/test/counts.json", func(w http.ResponseWriter, r *http.Request) {
		assertCountBody(t, r)
		setResponse(w)
	})

	params := &PremiumSearchCountTweetParams{
		Query:    "url:\"http://example.com\"",
		Tag:      "8HYG54ZGTU",
		FromDate: "201512220000",
		ToDate:   "201712220000",
		Bucket:   "day",
		Next:     "NTcxODIyMDMyODMwMjU1MTA0",
	}
	expected := &PremiumSearchCount{
		Results: []TweetCount{
			TweetCount{
				TimePeriod: "201701010000",
				Count:      32,
			},
			TweetCount{
				TimePeriod: "201701020000",
				Count:      45,
			},
		},
		TotalCount: 2027,
		RequestParameters: &RequestCountParameters{
			Bucket:   "day",
			FromDate: "201512200000",
			ToDate:   "201712200000",
		},
	}

	{
		client := NewClient(httpClient)
		search, _, err := client.PremiumSearch.CountFullArchive(
			params,
			"test",
		)
		assert.Nil(t, err)
		assert.Equal(t, expected, search)
	}
	{
		client := NewClient(httpClient)
		search, _, err := client.PremiumSearch.Count30Days(
			params,
			"test",
		)
		assert.Nil(t, err)
		assert.Equal(t, expected, search)
	}
}
