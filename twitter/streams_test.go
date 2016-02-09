package twitter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStream_MessageJSONError(t *testing.T) {
	badJSON := []byte(`{`)

	msg := getMessage(badJSON)
	assert.EqualError(t, msg.(error), "unexpected end of JSON input")
}

func TestStream_GetMessageTweet(t *testing.T) {
	// Example tweet json: https://raw.githubusercontent.com/ChimeraCoder/anaconda/master/json/statuses/show.json
	msgJSON := []byte(`{"created_at":"Tue Feb 19 08:04:41 +0000 2013","id":303777106620452864,"id_str":"303777106620452864","text":"golang-syd is in session. Dave Symonds is now talking about API design and protobufs. #golang http:\/\/t.co\/eSq3ROwu","source":"\u003ca href=\"http:\/\/twitter.com\/download\/android\" rel=\"nofollow\"\u003eTwitter for Android\u003c\/a\u003e","truncated":false,"in_reply_to_status_id":null,"in_reply_to_status_id_str":null,"in_reply_to_user_id":null,"in_reply_to_user_id_str":null,"in_reply_to_screen_name":null,"user":{"id":113419064,"id_str":"113419064","name":"Go","screen_name":"golang","location":"","description":"Go will make you love programming again. I promise.","url":"http:\/\/t.co\/C4svVTkUmj","entities":{"url":{"urls":[{"url":"http:\/\/t.co\/C4svVTkUmj","expanded_url":"http:\/\/golang.org\/","display_url":"golang.org","indices":[0,22]}]},"description":{"urls":[]}},"protected":false,"followers_count":34571,"friends_count":18,"listed_count":983,"created_at":"Thu Feb 11 18:04:38 +0000 2010","favourites_count":202,"utc_offset":-32400,"time_zone":"Alaska","geo_enabled":false,"verified":false,"statuses_count":1920,"lang":"en","contributors_enabled":false,"is_translator":false,"is_translation_enabled":false,"profile_background_color":"C0DEED","profile_background_image_url":"http:\/\/abs.twimg.com\/images\/themes\/theme1\/bg.png","profile_background_image_url_https":"https:\/\/abs.twimg.com\/images\/themes\/theme1\/bg.png","profile_background_tile":false,"profile_image_url":"http:\/\/pbs.twimg.com\/profile_images\/2388595262\/v02jhlxou71qagr6mwet_normal.png","profile_image_url_https":"https:\/\/pbs.twimg.com\/profile_images\/2388595262\/v02jhlxou71qagr6mwet_normal.png","profile_banner_url":"https:\/\/pbs.twimg.com\/profile_banners\/113419064\/1398369112","profile_link_color":"0084B4","profile_sidebar_border_color":"C0DEED","profile_sidebar_fill_color":"DDEEF6","profile_text_color":"333333","profile_use_background_image":true,"has_extended_profile":false,"default_profile":true,"default_profile_image":false,"following":true,"follow_request_sent":false,"notifications":false},"geo":null,"coordinates":null,"place":null,"contributors":null,"is_quote_status":false,"retweet_count":2,"favorite_count":3,"entities":{"hashtags":[{"text":"golang","indices":[86,93]}],"symbols":[],"user_mentions":[],"urls":[],"media":[{"id":303777106628841472,"id_str":"303777106628841472","indices":[94,114],"media_url":"http:\/\/pbs.twimg.com\/media\/BDc7q0OCEAAoe2C.jpg","media_url_https":"https:\/\/pbs.twimg.com\/media\/BDc7q0OCEAAoe2C.jpg","url":"http:\/\/t.co\/eSq3ROwu","display_url":"pic.twitter.com\/eSq3ROwu","expanded_url":"http:\/\/twitter.com\/golang\/status\/303777106620452864\/photo\/1","type":"photo","sizes":{"small":{"w":340,"h":255,"resize":"fit"},"medium":{"w":600,"h":450,"resize":"fit"},"thumb":{"w":150,"h":150,"resize":"crop"},"large":{"w":1024,"h":768,"resize":"fit"}}}]},"extended_entities":{"media":[{"id":303777106628841472,"id_str":"303777106628841472","indices":[94,114],"media_url":"http:\/\/pbs.twimg.com\/media\/BDc7q0OCEAAoe2C.jpg","media_url_https":"https:\/\/pbs.twimg.com\/media\/BDc7q0OCEAAoe2C.jpg","url":"http:\/\/t.co\/eSq3ROwu","display_url":"pic.twitter.com\/eSq3ROwu","expanded_url":"http:\/\/twitter.com\/golang\/status\/303777106620452864\/photo\/1","type":"photo","sizes":{"small":{"w":340,"h":255,"resize":"fit"},"medium":{"w":600,"h":450,"resize":"fit"},"thumb":{"w":150,"h":150,"resize":"crop"},"large":{"w":1024,"h":768,"resize":"fit"}}}]},"favorited":false,"retweeted":false,"possibly_sensitive":false,"possibly_sensitive_appealable":false,"lang":"en"}`)

	msg := getMessage(msgJSON)
	assert.IsType(t, &Tweet{}, msg)
}

func TestStream_GetMessageDirectMessage(t *testing.T) {
	// Example direct message: https://github.com/ChimeraCoder/anaconda/blob/master/json/direct_messages/new.json
	msgJSON := []byte(`{"direct_message": {"id":666024290140217347,"id_str":"666024290140217347","text":"Test the anaconda lib","sender":{"id":182675886,"id_str":"182675886","name":"Aditya Mukerjee","screen_name":"chimeracoder","location":"New York, NY","description":"Risk engineer at @stripe. Linux dev, statistician. Writing lots of Go. Alum of @recursecenter, @cornell_tech, @columbia","url":"http:\/\/t.co\/YhPyE6aJso","entities":{"url":{"urls":[{"url":"http:\/\/t.co\/YhPyE6aJso","expanded_url":"http:\/\/www.adityamukerjee.net","display_url":"adityamukerjee.net","indices":[0,22]}]},"description":{"urls":[]}},"protected":false,"followers_count":2872,"friends_count":769,"listed_count":160,"created_at":"Wed Aug 25 03:49:41 +0000 2010","favourites_count":2814,"utc_offset":-18000,"time_zone":"Eastern Time (US & Canada)","geo_enabled":false,"verified":false,"statuses_count":7798,"lang":"en","contributors_enabled":false,"is_translator":false,"is_translation_enabled":false,"profile_background_color":"C0DEED","profile_background_image_url":"http:\/\/abs.twimg.com\/images\/themes\/theme1\/bg.png","profile_background_image_url_https":"https:\/\/abs.twimg.com\/images\/themes\/theme1\/bg.png","profile_background_tile":false,"profile_image_url":"http:\/\/pbs.twimg.com\/profile_images\/1807988313\/230348_1870593437981_1035450059_32104665_3285049_n_cropped_normal.jpg","profile_image_url_https":"https:\/\/pbs.twimg.com\/profile_images\/1807988313\/230348_1870593437981_1035450059_32104665_3285049_n_cropped_normal.jpg","profile_link_color":"0084B4","profile_sidebar_border_color":"C0DEED","profile_sidebar_fill_color":"DDEEF6","profile_text_color":"333333","profile_use_background_image":true,"has_extended_profile":false,"default_profile":true,"default_profile_image":false,"following":false,"follow_request_sent":false,"notifications":false},"sender_id":182675886,"sender_id_str":"182675886","sender_screen_name":"chimeracoder","recipient":{"id":182675886,"id_str":"182675886","name":"Aditya Mukerjee","screen_name":"chimeracoder","location":"New York, NY","description":"Risk engineer at @stripe. Linux dev, statistician. Writing lots of Go. Alum of @recursecenter, @cornell_tech, @columbia","url":"http:\/\/t.co\/YhPyE6aJso","entities":{"url":{"urls":[{"url":"http:\/\/t.co\/YhPyE6aJso","expanded_url":"http:\/\/www.adityamukerjee.net","display_url":"adityamukerjee.net","indices":[0,22]}]},"description":{"urls":[]}},"protected":false,"followers_count":2872,"friends_count":769,"listed_count":160,"created_at":"Wed Aug 25 03:49:41 +0000 2010","favourites_count":2814,"utc_offset":-18000,"time_zone":"Eastern Time (US & Canada)","geo_enabled":false,"verified":false,"statuses_count":7798,"lang":"en","contributors_enabled":false,"is_translator":false,"is_translation_enabled":false,"profile_background_color":"C0DEED","profile_background_image_url":"http:\/\/abs.twimg.com\/images\/themes\/theme1\/bg.png","profile_background_image_url_https":"https:\/\/abs.twimg.com\/images\/themes\/theme1\/bg.png","profile_background_tile":false,"profile_image_url":"http:\/\/pbs.twimg.com\/profile_images\/1807988313\/230348_1870593437981_1035450059_32104665_3285049_n_cropped_normal.jpg","profile_image_url_https":"https:\/\/pbs.twimg.com\/profile_images\/1807988313\/230348_1870593437981_1035450059_32104665_3285049_n_cropped_normal.jpg","profile_link_color":"0084B4","profile_sidebar_border_color":"C0DEED","profile_sidebar_fill_color":"DDEEF6","profile_text_color":"333333","profile_use_background_image":true,"has_extended_profile":false,"default_profile":true,"default_profile_image":false,"following":false,"follow_request_sent":false,"notifications":false},"recipient_id":182675886,"recipient_id_str":"182675886","recipient_screen_name":"chimeracoder","created_at":"Sun Nov 15 22:45:39 +0000 2015","entities":{"hashtags":[],"symbols":[],"user_mentions":[],"urls":[]}} }`)

	msg := getMessage(msgJSON)
	assert.IsType(t, &DirectMessage{}, msg)
}

func TestStream_GetMessageDelete(t *testing.T) {
	msgJSON := []byte(`{"delete": { "id": 666024290140217347, "id_str": "666024290140217347", "user_id": 182675886, "user_id_str": "182675886" }}`)

	msg := getMessage(msgJSON)
	assert.IsType(t, &StatusDeletion{}, msg)
}

func TestStream_GetMessageLocationDeletion(t *testing.T) {
	msgJSON := []byte(`{"scrub_geo": { "up_to_status_id": 666024290140217347, "up_to_status_id_str": "666024290140217347", "user_id": 182675886, "user_id_str": "182675886" }}`)

	msg := getMessage(msgJSON)
	assert.IsType(t, &LocationDeletion{}, msg)
}

func TestStream_GetMessageStreamLimit(t *testing.T) {
	msgJSON := []byte(`{"limit": { "track": 10 }}`)

	msg := getMessage(msgJSON)
	assert.IsType(t, &StreamLimit{}, msg)
}

func TestStream_StatusWithheld(t *testing.T) {
	msgJSON := []byte(`{"status_withheld": { "id": 666024290140217347, "user_id": 182675886, "withheld_in_countries":["USA", "China"] }}`)

	msg := getMessage(msgJSON)
	assert.IsType(t, &StatusWithheld{}, msg)
}

func TestStream_UserWithheld(t *testing.T) {
	msgJSON := []byte(`{"user_withheld": { "id": 666024290140217347, "withheld_in_countries":["USA", "China"] }}`)

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

func TestStream_Unknown(t *testing.T) {
	msgJSON := []byte(`{"unknown_data": {"new_twitter_type":"unexpected"}}`)

	msg := getMessage(msgJSON)
	assert.IsType(t, map[string]interface{}{}, msg)
}

func TestStream_Stop(t *testing.T) {
	httpClient, _, server := testServer()
	defer server.Close()

	client := NewClient(httpClient)
	stream, err := client.Streams.User(nil)
	assert.NoError(t, err)
	stream.Stop()

}

func TestStream_User(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	reqCount := 0
	mux.HandleFunc("/1.1/user.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{}, r)
		switch reqCount {
		case 0:
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Transfer-Encoding", "chunked")
			fmt.Fprintf(w, `{"friends": [666024290140217347, 666024290140217349, 666024290140217342]}`+"\r\n"+"\r\n")
		default:
			// Only allow first request
			http.Error(w, "Stream API not available!", 130)
		}
		reqCount++
	})

	client := NewClient(httpClient)
	stream, err := client.Streams.User(nil)
	assert.NoError(t, err)
	defer stream.Stop()

	handled := testHandleStream(stream)

	assertReceive(t, handled, defaultTestTimeout, "stream expected to handle messages but timedout")
}

func TestStream_PublicSample(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	reqCount := 0
	mux.HandleFunc("/1.1/statuses/sample.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{}, r)
		switch reqCount {
		case 0:
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Transfer-Encoding", "chunked")
			fmt.Fprintf(w,
				`{"text": "Gophercon talks!"}`+"\r\n"+
					`{"text": "Gophercon super talks!"}`+"\r\n",
			)
		default:
			// Only allow first request
			http.Error(w, "Stream API not available!", 130)
		}
		reqCount++
	})

	client := NewClient(httpClient)
	stream, err := client.Streams.Sample(nil)
	assert.NoError(t, err)
	defer stream.Stop()

	handled := testHandleStream(stream)

	assertReceive(t, handled, defaultTestTimeout, "stream expected to handle messages but timedout")
	assertReceive(t, handled, defaultTestTimeout, "stream expected to handle messages but timedout")

}

func TestStream_PublicFilter(t *testing.T) {
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
			// Only allow first request
			http.Error(w, "Stream API not available!", 130)
		}
		reqCount++
	})

	client := NewClient(httpClient)
	stream, err := client.Streams.Filter(&StreamFilterParams{Track: []string{"gophercon", "golang"}})
	assert.NoError(t, err)
	defer stream.Stop()

	handled := testHandleStream(stream)

	assertReceive(t, handled, defaultTestTimeout, "stream expected to handle messages but timedout")
	assertReceive(t, handled, defaultTestTimeout, "stream expected to handle messages but timedout")

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
			// Only allow first request
			http.Error(w, "Stream API not available!", 130)
		}
		reqCount++
	})

	client := NewClient(httpClient)
	stream, err := client.Streams.Firehose(&StreamFirehoseParams{Count: 100})
	assert.NoError(t, err)
	defer stream.Stop()

	handled := testHandleStream(stream)

	assertReceive(t, handled, defaultTestTimeout, "stream expected to handle messages but timedout")
	assertReceive(t, handled, defaultTestTimeout, "stream expected to handle messages but timedout")
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
			// Only allow first request
			http.Error(w, "Stream API not available!", 130)
		}
		reqCount++
	})

	client := NewClient(httpClient)
	stream, err := client.Streams.Site(&StreamSiteParams{Follow: []string{"666024290140217347", "666024290140217349"}})
	assert.NoError(t, err)
	defer stream.Stop()

	handled := testHandleStream(stream)

	assertReceive(t, handled, defaultTestTimeout, "stream expected to handle messages but timedout")
	assertReceive(t, handled, defaultTestTimeout, "stream expected to handle messages but timedout")
}

func TestStream_Http503(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	reqCount := 0
	httpErrorSent := make(chan bool)
	mux.HandleFunc("/1.1/statuses/sample.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{}, r)
		switch reqCount {
		case 0:
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Transfer-Encoding", "chunked")
			fmt.Fprintf(w,
				`{"text": "Gophercon talks!"}`+"\r\n"+
					`{"text": "Gophercon super talks!"}`+"\r\n",
			)
		case 1:
			// Exponential backoff
			http.Error(w, "Service Unavailable", 503)
			httpErrorSent <- true
		default:
			// Only allow first 2 requests
			http.Error(w, "Unknown", 404)
		}
		reqCount++
	})

	client := NewClient(httpClient)
	stream, err := client.Streams.Sample(nil)
	assert.NoError(t, err)
	defer stream.Stop()

	handled := testHandleStream(stream)

	assertReceive(t, handled, defaultTestTimeout, "stream expected to handle messages but timedout")
	assertReceive(t, handled, defaultTestTimeout, "stream expected to handle messages but timedout")

	assertReceive(t, httpErrorSent, defaultTestTimeout, "stream expected to receive error")
}

func TestStream_HttpBackoff420(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	reqCount := 0
	httpErrorSent := make(chan bool)
	mux.HandleFunc("/1.1/statuses/sample.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{}, r)
		switch reqCount {
		case 0:
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Transfer-Encoding", "chunked")
			fmt.Fprintf(w,
				`{"text": "Gophercon talks!"}`+"\r\n"+
					`{"text": "Gophercon super talks!"}`+"\r\n",
			)
		case 1:
			// aggressive Exponential backoff
			http.Error(w, "Rate Limited", 420)
			httpErrorSent <- true
		default:
			// Only allow first 2 requests
			http.Error(w, "Unknown", 404)
		}
		reqCount++
	})

	client := NewClient(httpClient)
	stream, err := client.Streams.Sample(nil)
	assert.NoError(t, err)
	defer stream.Stop()

	handled := testHandleStream(stream)

	assertReceive(t, handled, defaultTestTimeout, "stream expected to handle messages but timedout")
	assertReceive(t, handled, defaultTestTimeout, "stream expected to handle messages but timedout")

	assertReceive(t, httpErrorSent, defaultTestTimeout, "stream expected to receive error")
}

func TestStream_HttpError404(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	reqCount := 0
	httpErrorSent := make(chan bool)
	mux.HandleFunc("/1.1/statuses/sample.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{}, r)
		switch reqCount {
		case 0:
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Transfer-Encoding", "chunked")
			fmt.Fprintf(w,
				`{"text": "Gophercon talks!"}`+"\r\n"+
					`{"text": "Gophercon super talks!"}`+"\r\n",
			)
		case 1:
			// disconnect no reconnect
			http.Error(w, "Unknown", 404)
			httpErrorSent <- true
		default:
			// Only allow first 2 requests
			http.Error(w, "Unknown", 404)
		}
		reqCount++
	})

	client := NewClient(httpClient)
	stream, err := client.Streams.Sample(nil)
	assert.NoError(t, err)
	defer stream.Stop()

	handled := testHandleStream(stream)

	assertReceive(t, handled, defaultTestTimeout, "stream expected to handle messages but timedout")
	assertReceive(t, handled, defaultTestTimeout, "stream expected to handle messages but timedout")

	assertReceive(t, httpErrorSent, defaultTestTimeout, "stream expected to receive error")
}
