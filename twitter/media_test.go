package twitter

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fake a server to handle the upload requests. Failure triggers:
// START:
//    11 byte images fail
// APPEND:
//    17 byte images fail
// FINALIZE
//    23 byte images fail

func uploadResponseFunc(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
	}
	w.Header().Set("Content-Type", "application/json")

	rt := r.FormValue("command")
	log.Printf("command is %q", rt)
	// No command means it's a bad request
	if rt == "" {
		log.Printf("no command")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch rt {
	case "INIT":
		tb := r.FormValue("total_bytes")
		// 11 byte requests trigger the
		if tb == "11" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("tb is %q", tb)
		fmt.Fprintf(w, `{"media_id": %v, "media_id_string": "%v", "size": %v, "expires_after_secs": 86400}`, tb, tb, tb)
	case "APPEND":
		mid := r.FormValue("media_id")
		if mid == "17" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case "FINALIZE":
		mid := r.FormValue("media_id")
		if mid == "23" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, `{"media_id": %v, "media_id_string": "%v", "size": %v, "expires_after_secs": 86400}`, mid, mid, mid)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"status":"bad"}`)
	}
}

func TestMediaService_Upload(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		filetype string
		wantErr  bool
		want     *MediaUploadResult
	}{
		{
			name:     "small OK",
			data:     []byte{1, 2, 3},
			filetype: "image/jpeg",
			want: &MediaUploadResult{
				MediaID:          3,
				MediaIDString:    "3",
				Size:             3,
				ExpiresAfterSecs: 86400,
			},
		},
		{
			name:     "multipart OK",
			data:     bytes.Repeat([]byte{50}, chunkSize*4),
			filetype: "video/mp4",
			want: &MediaUploadResult{
				MediaID:          chunkSize * 4,
				MediaIDString:    fmt.Sprintf("%v", chunkSize*4),
				Size:             chunkSize * 4,
				ExpiresAfterSecs: 86400,
			},
		},
		{
			name:     "start fails",
			data:     []byte{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
			filetype: "video/mp4",
			wantErr:  true,
		},
		{
			name:     "big fail",
			data:     bytes.Repeat([]byte{10}, maxSize+20),
			filetype: "image/gif",
			wantErr:  true,
		},
		{
			name:     "append fails",
			data:     bytes.Repeat([]byte{17}, 17),
			filetype: "image/jpeg",
			wantErr:  true,
		},
		{
			name:     "finalize fails",
			data:     bytes.Repeat([]byte{23}, 23),
			filetype: "image/gif",
			wantErr:  true,
		},
	}

	httpClient, mux, server := testServer()
	defer server.Close()
	mux.HandleFunc("/1.1/media/upload.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		uploadResponseFunc(w, r)
	})

	client := NewClient(httpClient)

	for _, test := range tests {
		resp, _, err := client.Media.Upload(test.data, test.filetype)
		if err != nil {
			if !test.wantErr {
				t.Errorf("Media.Upload(%v): err: %v", test.name, err)
			}
			continue
		}
		if err == nil {
			assert.Equal(t, test.want, resp)
		}

	}

}

func TestMediaService_Status(t *testing.T) {

	httpClient, mux, server := testServer()
	defer server.Close()
	mux.HandleFunc("/1.1/media/upload.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"media_id": 123, "media_id_string": "123", "expires_after_secs": 86400, "processing_info": {"state": "succeeded", "progress_percent": 100}}`)
	})

	expected := &MediaStatusResult{
		MediaID:          123,
		MediaIDString:    "123",
		ExpiresAfterSecs: 86400,
		ProcessingInfo: &MediaProcessingInfo{
			State:           "succeeded",
			ProgressPercent: 100,
		},
	}

	client := NewClient(httpClient)
	result, _, err := client.Media.Status(123)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}
