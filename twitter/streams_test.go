package twitter

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStream_MessageJSONError(t *testing.T) {
	badJSON := []byte(`{`)
	msg := getMessage(badJSON)
	assert.EqualError(t, msg.(error), "unexpected end of JSON input")
}

func TestStream_GetMessageTweet(t *testing.T) {
	msgJSON := []byte(`{"id": 20, "text": "just setting up my twttr", "retweet_count": "68535"}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, &Tweet{}, msg)
}

func TestStream_GetMessageDirectMessage(t *testing.T) {
	msgJSON := []byte(`{"direct_message": {"id": 666024290140217347}}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, &DirectMessage{}, msg)
}

func TestStream_GetMessageDelete(t *testing.T) {
	msgJSON := []byte(`{"delete": { "id": 20}}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, &StatusDeletion{}, msg)
}

func TestStream_GetMessageLocationDeletion(t *testing.T) {
	msgJSON := []byte(`{"scrub_geo": { "up_to_status_id": 20}}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, &LocationDeletion{}, msg)
}

func TestStream_GetMessageStreamLimit(t *testing.T) {
	msgJSON := []byte(`{"limit": { "track": 10 }}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, &StreamLimit{}, msg)
}

func TestStream_StatusWithheld(t *testing.T) {
	msgJSON := []byte(`{"status_withheld": { "id": 20, "user_id": 12, "withheld_in_countries":["USA", "China"] }}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, &StatusWithheld{}, msg)
}

func TestStream_UserWithheld(t *testing.T) {
	msgJSON := []byte(`{"user_withheld": { "id": 12, "withheld_in_countries":["USA", "China"] }}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, &UserWithheld{}, msg)
}

func TestStream_StreamDisconnect(t *testing.T) {
	msgJSON := []byte(`{"disconnect": { "code": "420", "stream_name": "streaming stuff", "reason": "too many connections" }}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, &StreamDisconnect{}, msg)
}

func TestStream_StallWarning(t *testing.T) {
	msgJSON := []byte(`{"warning": { "code": "420", "percent_full": 90, "message": "a lot of messages" }}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, &StallWarning{}, msg)
}

func TestStream_FriendsList(t *testing.T) {
	msgJSON := []byte(`{"friends": [666024290140217347, 666024290140217349, 666024290140217342]}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, &FriendsList{}, msg)
}

func TestStream_Event(t *testing.T) {
	msgJSON := []byte(`{"event": "block", "target": {"name": "XKCD Comic", "favourites_count": 2}, "source": {"name": "XKCD Comic2", "favourites_count": 3}, "created_at": "Sat Sep 4 16:10:54 +0000 2010"}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, &Event{}, msg)
}

func TestStream_APIError(t *testing.T) {
	msgJSON := []byte(`{"errors":[{"code":215,"message":"Bad Authentication data."}]}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, &APIError{}, msg)
}

func TestStream_Unknown(t *testing.T) {
	msgJSON := []byte(`{"unknown_data": {"new_twitter_type":"unexpected"}}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, map[string]interface{}{}, msg)
}

func TestStream_Filter(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	reqCount := 0
	mux.HandleFunc("/1.1/statuses/filter.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertQuery(t, map[string]string{"track": "gophercon,golang"}, r)
		switch reqCount {
		case 0:
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Transfer-Encoding", "chunked")
			fmt.Fprintf(w,
				`{"text": "Gophercon talks!"}`+"\r\n"+
					`{"text": "Gophercon super talks!"}`+"\r\n",
			)
		default:
			// Simulate stream disconnect
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		}
		reqCount++
	})

	counts := &counter{}
	demux := newCounterDemux(counts)
	client := NewClient(httpClient)
	streamFilterParams := &StreamFilterParams{
		Track: []string{"gophercon", "golang"},
	}
	client.Streams.exponentialBackoff = &BackOffRecorder{MaxRetries: 1}
	client.Streams.aggressiveExponentialBackoff = &BackOffRecorder{MaxRetries: 1}
	stream, err := client.Streams.Filter(streamFilterParams)
	// assert that the expected messages are received
	assert.NoError(t, err)
	defer stream.Stop()
	for message := range stream.Messages {
		demux.Handle(message)
	}
	expectedCounts := &counter{all: 3, other: 3}
	assert.Equal(t, expectedCounts, counts)
}

func TestStream_Sample(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	reqCount := 0
	mux.HandleFunc("/1.1/statuses/sample.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"stall_warnings": "true"}, r)
		switch reqCount {
		case 0:
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Transfer-Encoding", "chunked")
			fmt.Fprintf(w,
				`{"text": "Gophercon talks!"}`+"\r\n"+
					`{"text": "Gophercon super talks!"}`+"\r\n",
			)
		default:
			// Simulate stream disconnect
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		}
		reqCount++
	})

	counts := &counter{}
	demux := newCounterDemux(counts)
	client := NewClient(httpClient)
	streamSampleParams := &StreamSampleParams{
		StallWarnings: Bool(true),
	}
	client.Streams.exponentialBackoff = &BackOffRecorder{MaxRetries: 1}
	client.Streams.aggressiveExponentialBackoff = &BackOffRecorder{MaxRetries: 1}
	stream, err := client.Streams.Sample(streamSampleParams)
	// assert that the expected messages are received
	assert.NoError(t, err)
	defer stream.Stop()
	for message := range stream.Messages {
		demux.Handle(message)
	}
	expectedCounts := &counter{all: 3, other: 3}
	assert.Equal(t, expectedCounts, counts)
}

func TestStream_User(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	reqCount := 0
	mux.HandleFunc("/1.1/user.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"stall_warnings": "true", "with": "followings"}, r)
		switch reqCount {
		case 0:
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Transfer-Encoding", "chunked")
			fmt.Fprintf(w, `{"friends": [666024290140217347, 666024290140217349, 666024290140217342]}`+"\r\n"+"\r\n")
		default:
			// Simulate stream disconnect
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		}
		reqCount++
	})

	counts := &counter{}
	demux := newCounterDemux(counts)
	client := NewClient(httpClient)
	streamUserParams := &StreamUserParams{
		StallWarnings: Bool(true),
		With:          "followings",
	}
	client.Streams.exponentialBackoff = &BackOffRecorder{MaxRetries: 1}
	client.Streams.aggressiveExponentialBackoff = &BackOffRecorder{MaxRetries: 1}
	stream, err := client.Streams.User(streamUserParams)
	// assert that the expected messages are received
	assert.NoError(t, err)
	defer stream.Stop()
	for message := range stream.Messages {
		demux.Handle(message)
	}
	expectedCounts := &counter{all: 2, friendsList: 1, other: 1}
	assert.Equal(t, expectedCounts, counts)
}

func TestStream_User_TooManyFriends(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	reqCount := 0
	mux.HandleFunc("/1.1/user.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"stall_warnings": "true", "with": "followings"}, r)
		switch reqCount {
		case 0:
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Transfer-Encoding", "chunked")
			// The first friend list message is more than bufio.MaxScanTokenSize (65536) bytes
			friendsList := "[" + strings.Repeat("1234567890, ", 7000) + "1234567890]"
			fmt.Fprintf(w, `{"friends": %s}`+"\r\n"+"\r\n", friendsList)
		default:
			// Simulate stream disconnect
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		}
		reqCount++
	})

	counts := &counter{}
	demux := newCounterDemux(counts)
	client := NewClient(httpClient)
	streamUserParams := &StreamUserParams{
		StallWarnings: Bool(true),
		With:          "followings",
	}
	client.Streams.exponentialBackoff = &BackOffRecorder{MaxRetries: 1}
	client.Streams.aggressiveExponentialBackoff = &BackOffRecorder{MaxRetries: 1}
	stream, err := client.Streams.User(streamUserParams)
	// assert that the expected messages are received
	assert.NoError(t, err)
	defer stream.Stop()
	for message := range stream.Messages {
		demux.Handle(message)
	}
	expectedCounts := &counter{all: 2, friendsList: 1, other: 1}
	assert.Equal(t, expectedCounts, counts)
}

func TestStream_Site(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	reqCount := 0
	mux.HandleFunc("/1.1/site.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"follow": "666024290140217347,666024290140217349"}, r)
		switch reqCount {
		case 0:
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Transfer-Encoding", "chunked")
			fmt.Fprintf(w,
				`{"text": "Gophercon talks!"}`+"\r\n"+
					`{"text": "Gophercon super talks!"}`+"\r\n",
			)
		default:
			// Simulate stream disconnect
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		}
		reqCount++
	})

	counts := &counter{}
	demux := newCounterDemux(counts)
	client := NewClient(httpClient)
	streamSiteParams := &StreamSiteParams{
		Follow: []string{"666024290140217347", "666024290140217349"},
	}
	client.Streams.exponentialBackoff = &BackOffRecorder{MaxRetries: 1}
	client.Streams.aggressiveExponentialBackoff = &BackOffRecorder{MaxRetries: 1}
	stream, err := client.Streams.Site(streamSiteParams)
	// assert that the expected messages are received
	assert.NoError(t, err)
	for message := range stream.Messages {
		demux.Handle(message)
	}
	expectedCounts := &counter{all: 3, other: 3}
	assert.Equal(t, expectedCounts, counts)
}

func TestStream_PublicFirehose(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	reqCount := 0
	mux.HandleFunc("/1.1/statuses/firehose.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"count": "100"}, r)
		switch reqCount {
		case 0:
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Transfer-Encoding", "chunked")
			fmt.Fprintf(w,
				`{"text": "Gophercon talks!"}`+"\r\n"+
					`{"text": "Gophercon super talks!"}`+"\r\n",
			)
		default:
			// Simulate stream disconnect
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		}
		reqCount++
	})

	counts := &counter{}
	demux := newCounterDemux(counts)
	client := NewClient(httpClient)
	streamFirehoseParams := &StreamFirehoseParams{
		Count: 100,
	}
	client.Streams.exponentialBackoff = &BackOffRecorder{MaxRetries: 1}
	client.Streams.aggressiveExponentialBackoff = &BackOffRecorder{MaxRetries: 1}
	stream, err := client.Streams.Firehose(streamFirehoseParams)
	// assert that the expected messages are received
	assert.NoError(t, err)
	defer stream.Stop()
	for message := range stream.Messages {
		demux.Handle(message)
	}
	expectedCounts := &counter{all: 3, other: 3}
	assert.Equal(t, expectedCounts, counts)
}

func TestStreamRetry_ExponentialBackoff(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
	})

	stream := &Stream{
		client:   httpClient,
		Messages: make(chan interface{}),
		done:     make(chan struct{}),
		group:    &sync.WaitGroup{},
	}
	stream.group.Add(1)
	req, _ := http.NewRequest("GET", "http://example.com/", nil)
	expBackoff := &BackOffRecorder{MaxRetries: 1}
	// receive messages and count types, stop receiving after max retries
	counts := &counter{}
	demux := newCounterDemux(counts)

	go stream.retry(req, expBackoff, nil)
	defer stream.Stop()
	for message := range stream.Messages {
		demux.Handle(message)
	}

	// assert exponential backoff in response to 503
	assert.Equal(t, 1, expBackoff.Count)
	expectedCounts := &counter{all: 1, other: 1}
	assert.Equal(t, expectedCounts, counts)
}

func TestStreamRetry_AggressiveBackoff(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	reqCount := 0
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch reqCount {
		case 0:
			http.Error(w, "Enhance your calm", 420)
		default:
			w.WriteHeader(http.StatusTooManyRequests)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{"errors": [{"code": 88, "message": "Rate limit exceeded"}]}`)
		}
		reqCount++
	})
	stream := &Stream{
		client:   httpClient,
		Messages: make(chan interface{}),
		done:     make(chan struct{}),
		group:    &sync.WaitGroup{},
	}
	stream.group.Add(1)
	req, _ := http.NewRequest("GET", "http://example.com/", nil)
	aggExpBackoff := &BackOffRecorder{MaxRetries: 2}

	// receive messages and count types, stop receiving after max retries
	counts := &counter{}
	demux := newCounterDemux(counts)

	go stream.retry(req, nil, aggExpBackoff)
	defer stream.Stop()

	for message := range stream.Messages {
		demux.Handle(message)
	}

	// assert aggressive exponential backoff in response to 420 and 429
	assert.Equal(t, 2, aggExpBackoff.Count)
	expectedCounts := &counter{all: 2, apiError: 1, other: 1}
	assert.Equal(t, expectedCounts, counts)
}
