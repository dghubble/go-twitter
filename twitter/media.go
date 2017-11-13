package twitter

import (
	"encoding/base64"
	"github.com/dghubble/sling"
	"io/ioutil"
	"net/http"
)

const twitterUploadAPI = "https://upload.twitter.com/1.1/"

type Media struct {
	MediaId             int64       `json:"media_id"`
	MediaIdString       string      `json:"media_id_string"`
	Size                int32       `json:"size"`
	ExpiresAfterSeconds int32       `json:"expires_after_secs"`
	Image               *MediaImage `json:"image"`
}

type MediaImage struct {
	ImageType string `json:"image_type"`
	Width     int32  `json:"w"`
	Height    int32  `json:"h"`
}

type MediaParams struct {
	MediaData string `url:"media_data"`
}

type MediaService struct {
	sling *sling.Sling
}

func newMediaService(sling *sling.Sling) *MediaService {
	return &MediaService{
		sling: sling.Path("media/"),
	}
}

func (m *MediaService) Encode(path string) (string, error) {
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buff), nil
}

func (m *MediaService) UploadFile(filePath string) (*Media, *http.Response, error) {
	EncodedImage, _ := m.Encode(filePath)
	apiError := new(APIError)
	mediaImage := new(Media)

	params := MediaParams{MediaData: EncodedImage}
	resp, err := m.sling.New().Base(twitterUploadAPI).Path("media/").Post("upload.json").BodyForm(params).Receive(mediaImage, apiError)

	return mediaImage, resp, relevantError(err, *apiError)

}
