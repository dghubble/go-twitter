package twitter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigService_Get(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/help/configuration.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, ` {"characters_reserved_per_media": 24, "dm_text_character_limit": 10000, "max_media_per_upload": 1, "photo_size_limit": 3145728, "photo_sizes": { "large": { "h": 2048, "resize": "fit", "w": 1024 }, "medium": { "h": 1200, "resize": "fit", "w": 600 }, "small": { "h": 480, "resize": "fit", "w": 340 }, "thumb": { "h": 150, "resize": "crop", "w": 150 } }, "short_url_length": 23, "short_url_length_https": 23, "non_username_paths": [ "about" ] }`)
	})

	client := NewClient(httpClient)
	status, _, err := client.Config.Get()
	expected := &Config{CharactersReservedPerMedia: 24, DMTextCharacterLimit: 10000, MaxMediaPerUpload: 1, PhotoSizeLimit: 3145728, PhotoSizes: &PhotoSizes{Large: &SinglePhotoSize{Height: 2048, Width: 1024, Resize: "fit"}, Medium: &SinglePhotoSize{Height: 1200, Width: 600, Resize: "fit"}, Small: &SinglePhotoSize{Height: 480, Width: 340, Resize: "fit"}, Thumb: &SinglePhotoSize{Height: 150, Width: 150, Resize: "crop"}}, ShortURLLength: 23, ShortURLLengthHTTPS: 23, NonUsernamePaths: []string{"about"}}
	assert.Nil(t, err)
	assert.Equal(t, expected, status)
}
