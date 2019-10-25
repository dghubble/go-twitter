package twitter

import (
	b64 "encoding/base64"
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

//MediaService ... provides methods for accessing Twitter status API endpoints.
type MediaService struct {
	sling *sling.Sling
}

// newStatusService returns a new StatusService.
func newMediaService(sling *sling.Sling) *MediaService {
	return &MediaService{
		sling: sling.Path("media/"),
	}
}

//MediaUploadParams ... are the parameters for StatusService.Update
type MediaUploadParams struct {
	File     []byte
	MimeType string
}

//Media ... response of uploaded file
type Media struct {
	MediaID          int64  `json:"media_id"`
	MediaIDString    string `json:"media_id_string"`
	ExpiresAfterSecs uint64 `json:"expires_after_secs"`
}

type mediaUploadCommand struct {
	Command      string `url:"command,omitempty"`
	MediaID      string `url:"media_id,omitempty"`
	MediaType    string `url:"media_type,omitempty"`
	MediaData    string `url:"media_data,omitempty"`
	SegmentIndex string `url:"segment_index,omitempty"`
	TotalBytes   string `url:"total_bytes,omitempty"`
}

func (m MediaUploadParams) getTotalBytes() int {
	if m.File != nil {
		return len(m.File)
	}

	return 0
}

// Upload media file
// Requires a user auth context.
// https://dev.twitter.com/rest/reference/post/media/upload
func (s *MediaService) Upload(params *MediaUploadParams) (*Media, *http.Response, error) {
	var resp *http.Response
	var err error
	var twitterMediaID *Media

	twitterMediaID, resp, err = s.mediaInit(params)
	if err != nil {
		return nil, resp, err
	}

	resp, err = s.mediaAppend(twitterMediaID, params)
	if err != nil {
		return nil, resp, err
	}

	resp, err = s.mediaFinilize(twitterMediaID.MediaID)
	if err != nil {
		return nil, resp, err
	}

	return twitterMediaID, resp, nil
}

func (s *MediaService) mediaInit(p *MediaUploadParams) (*Media, *http.Response, error) {
	paramsBody := mediaUploadCommand{
		Command:    "INIT",
		MediaType:  p.MimeType,
		TotalBytes: fmt.Sprintf("%d", p.getTotalBytes()),
	}

	twitterMediaID := new(Media)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("upload.json").Add("Content-Type", "application/x-www-form-urlencoded").BodyForm(paramsBody).Receive(twitterMediaID, apiError)
	return twitterMediaID, resp, relevantError(err, *apiError)
}

func (s *MediaService) mediaAppend(twitterMediaID *Media, params *MediaUploadParams) (*http.Response, error) {
	media := params.File
	mediaID := twitterMediaID.MediaIDString
	mediaBase64 := b64.StdEncoding.EncodeToString(media)

	step := 500 * 1024
	for i := 0; i*step < len(mediaBase64); i++ {
		rangeBegining := i * step
		rangeEnd := (i + 1) * step
		if rangeEnd > len(mediaBase64) {
			rangeEnd = len(mediaBase64)
		}
		_ = rangeBegining
		params := mediaUploadCommand{
			Command:      "APPEND",
			MediaID:      mediaID,
			MediaData:    mediaBase64[rangeBegining:rangeEnd],
			SegmentIndex: fmt.Sprint(i),
		}

		apiError := new(APIError)
		resp, err := s.sling.New().Post("upload.json").Add("Content-Type", "application/x-www-form-urlencoded").BodyForm(params).Receive(nil, apiError)
		if err != nil {
			return resp, relevantError(err, *apiError)
		}
	}

	return nil, nil
}

func (s *MediaService) mediaFinilize(mediaID int64) (*http.Response, error) {
	params := mediaUploadCommand{
		Command: "FINALIZE",
		MediaID: fmt.Sprint(mediaID),
	}

	apiError := new(APIError)
	resp, err := s.sling.New().Post("upload.json").Add("Content-Type", "application/x-www-form-urlencoded").BodyForm(params).Receive(nil, apiError)
	if err != nil {
		return resp, relevantError(err, *apiError)
	}

	return resp, nil
}
