package twitter

import (
	"encoding/base64"
	"github.com/dghubble/sling"
	"io/ioutil"
	"net/http"
)

const twitterUploadAPI = "https://upload.twitter.com/1.1/"

// Media is the structure received from Twitter after upload
type Media struct {
	MediaID             int64       `json:"media_id"`
	MediaIDString       string      `json:"media_id_string"`
	Size                int32       `json:"size"`
	ExpiresAfterSeconds int32       `json:"expires_after_secs"`
	Image               *MediaImage `json:"image"`
}

//MediaImage is the structure within Media that contains image information
type MediaImage struct {
	ImageType string `json:"image_type"`
	Width     int32  `json:"w"`
	Height    int32  `json:"h"`
}

// MediaParams is the struct expected by Twitter for uploading
type MediaParams struct {
	MediaData string `url:"media_data"`
}

// MediaService provides methods for accessing Twitter upload API endpoints.
type MediaService struct {
	sling *sling.Sling
}

func newMediaService(sling *sling.Sling) *MediaService {
	return &MediaService{
		sling: sling.Path("media/"),
	}
}

// Encode transfers the data of the image into base64
func (m *MediaService) Encode(path string) (string, error) {
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buff), nil
}

// UploadFile is a uploader for an image to twitter
// https://developer.twitter.com/en/docs/media/upload-media/overview
func (m *MediaService) UploadFile(filePath string) (*Media, *http.Response, error) {
	EncodedImage, _ := m.Encode(filePath)
	apiError := new(APIError)
	mediaImage := new(Media)

	params := MediaParams{MediaData: EncodedImage}
	resp, err := m.sling.New().Base(twitterUploadAPI).Path("media/").Post("upload.json").BodyForm(params).Receive(mediaImage, apiError)

	return mediaImage, resp, relevantError(err, *apiError)

}
