package twitter

import (
	"encoding/base64"
	"fmt"
	"github.com/dghubble/sling"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const twitterUploadAPI string = "https://upload.twitter.com/1.1/"
const twitterChunkSize int64 = 1048576
const twitterMaxSize int64 = 15728640
const twitterCommandAppend = "APPEND"
const twitterCommandInit = "INIT"
const twitterCommandFinalize = "FINALIZE"
const twitterCommandStatus = "STATUS"

// Media is the structure received from Twitter after upload
type Media struct {
	MediaID             int64       `json:"media_id"`
	MediaIDString       string      `json:"media_id_string"`
	Size                int32       `json:"size"`
	ExpiresAfterSeconds int32       `json:"expires_after_secs"`
	Image               *MediaImage `json:"image"`
}

// MediaImage is the structure within Media that contains image information
type MediaImage struct {
	ImageType string `json:"image_type"`
	Width     int32  `json:"w"`
	Height    int32  `json:"h"`
}

// MediaVideo is structure for media Video
type MediaVideo struct {
	Video int64 `json:"video_type"`
}

// MediaInit is the structure to init an chunked upload
type MediaInit struct {
	MediaID             int64  `json:"media_id"`
	MediaIDString       string `json:"media_id_string"`
	ExpiresAfterSeconds int32  `json:"expires_after_secs"`
	TotalBytes          int64
	FilePath            string
}

// MediaInitParam Params for chunked upload init
type MediaInitParam struct {
	Command          string `url:"command"`
	TotalBytes       int64  `url:"total_bytes,omitempty"`
	MediaType        string `url:"media_type,omitempty"`
	MediaCategory    string `url:"media_category,omitempty"`
	AdditionalOwners string `url:"additional_owners,omitempty"`
}

// MediaAppend is an empty result set
type MediaAppend struct {
}

// MediaFinalizeParam Params for finalizing
type MediaFinalizeParam struct {
	Command string `url:"command"`
	MediaID int64  `url:"media_id"`
}

// MediaAppendParam Params for appending chunk of image
type MediaAppendParam struct {
	Command      string `url:"command"`
	MediaID      int64  `url:"media_id"`
	SegmentIndex int16  `url:"segment_index"`
	MediaData    string `url:"media_data"`
}

// MediaFinalize Response struct Finalize
type MediaFinalize struct {
	MediaID             int64       `json:"media_id"`
	MediaIDString       string      `json:"media_id_string"`
	Size                int64       `json:"size"`
	ExpiresAfterSeconds int32       `json:"expires_after_secs"`
	Image               *MediaImage `json:"image"`
	Video               *MediaVideo `json:"video"`
}

// MediaStatusParam Status Param
type MediaStatusParam struct {
	Command string `url:"command"`
	MediaID int64  `url:"media_id"`
}

// MediaStatus Result set
type MediaStatus struct {
	MediaID        int64                `json:"media_id"`
	MediaIDString  string               `json:"media_id_string"`
	ProcessingInfo *MediaProcessingInfo `json:"processing_info"`
}

// MediaProcessingInfo information Result set
type MediaProcessingInfo struct {
	Stage           string                    `json:"state"`
	CheckAfterSecs  int                       `json:"check_after_secs"`
	ProgressPercent int                       `json:"progress_percent"`
	Error           *MediaProcessingInfoError `json:"error"`
}

// MediaProcessingInfoError Processing error
type MediaProcessingInfoError struct {
	Code    int    `json:"code"`
	Name    string `json:"name"`
	Message string `json:"message"`
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

// EncodeFile transfers the data of the image into base64
func (m *MediaService) EncodeFile(path string) (string, error) {
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buff), nil
}

// EncodeFileChunk transfers a chunk of the image into base64
func (m *MediaService) EncodeFileChunk(path string, segementIndex int16) (string, error) {
	var startPoint int64
	var chunkSize int64

	inputFile, err := os.Open(path)
	if err != nil {
		return "", err
	}

	fileInfo, err := inputFile.Stat()
	if err != nil {
		return "", err
	}

	chunkSize = twitterChunkSize
	if fileInfo.Size() < chunkSize {
		chunkSize = int64(fileInfo.Size())
	} else {
		startPoint = (int64(segementIndex) * twitterChunkSize)
		if (startPoint + twitterChunkSize) > fileInfo.Size() {
			chunkSize = int64(fileInfo.Size() % twitterChunkSize)
		} else {
			chunkSize = twitterChunkSize
		}
	}

	inputFile.Seek(startPoint, 0)
	chunk := make([]byte, chunkSize)
	inputFile.Read(chunk)
	return base64.StdEncoding.EncodeToString(chunk), nil
}

// UploadFile is a uploader for an image to twitter
// https://developer.twitter.com/en/docs/media/upload-media/overview
func (m *MediaService) UploadFile(filePath string) (*Media, *http.Response, error) {
	EncodedImage, _ := m.EncodeFile(filePath)
	apiError := new(APIError)
	mediaImage := new(Media)

	_, fileError := os.Stat(filePath)
	if fileError != nil {
		log.Fatal(fmt.Sprintf("File not found at location: %s", filePath))
	}

	params := MediaParams{MediaData: EncodedImage}
	resp, err := m.sling.New().Base(twitterUploadAPI).Path("media/").Post("upload.json").BodyForm(params).Receive(mediaImage, apiError)

	return mediaImage, resp, relevantError(err, *apiError)

}

// UploadInit initiate chunked upload
// https://developer.twitter.com/en/docs/media/upload-media/overview
func (m *MediaService) UploadInit(filePath string, mediaType string, category string, additionalOwners []string) (*MediaInit, *http.Response, error) {

	fileInfo, fileError := os.Stat(filePath)
	if fileError != nil {
		log.Fatal(fmt.Sprintf("File not found at location: %s", filePath))
	}

	if fileInfo.Size() > twitterMaxSize {
		log.Fatal("File is larger than the allowed max upload size")
	}

	apiError := new(APIError)
	mediaInit := new(MediaInit)

	params := MediaInitParam{
		Command:          twitterCommandInit,
		TotalBytes:       fileInfo.Size(),
		MediaType:        mediaType,
		MediaCategory:    category,
		AdditionalOwners: strings.Join(additionalOwners, ","),
	}

	resp, err := m.sling.New().Base(twitterUploadAPI).Path("media/").Post("upload.json").QueryStruct(params).Receive(mediaInit, apiError)

	mediaInit.TotalBytes = fileInfo.Size()
	mediaInit.FilePath = filePath

	return mediaInit, resp, relevantError(err, *apiError)
}

// UploadAppendChunks Upload the file chunks
// https://developer.twitter.com/en/docs/media/upload-media/api-reference/post-media-upload-append
func (m *MediaService) UploadAppendChunks(init *MediaInit) *http.Response {

	params := MediaAppendParam{
		Command: twitterCommandAppend,
		MediaID: init.MediaID,
	}
	var chunksNumber = int(init.TotalBytes / twitterChunkSize)

	for segmentIndex := 0; segmentIndex <= chunksNumber; segmentIndex++ {
		base64Chunk, _ := m.EncodeFileChunk(init.FilePath, int16(segmentIndex))
		params.SegmentIndex = int16(segmentIndex)
		params.MediaData = base64Chunk

		resp, _ := m.UploadAppend(params)
		if resp.StatusCode < 200 && resp.StatusCode > 299 {
			return resp
		}
	}
	return nil
}

// UploadAppend append chunk to uploaded file
// https://developer.twitter.com/en/docs/media/upload-media/api-reference/post-media-upload-append
func (m *MediaService) UploadAppend(appendParam MediaAppendParam) (*http.Response, error) {

	apiError := new(APIError)
	mediaAppend := new(MediaAppend)

	resp, err := m.sling.New().Base(twitterUploadAPI).Path("media/").Post("upload.json").BodyForm(appendParam).Receive(mediaAppend, apiError)
	return resp, relevantError(err, *apiError)
}

// UploadFinalize Finalize the upload
// https://developer.twitter.com/en/docs/media/upload-media/api-reference/post-media-upload-finalize
func (m *MediaService) UploadFinalize(init *MediaInit) (*MediaFinalize, *http.Response, error) {
	params := MediaFinalizeParam{
		Command: twitterCommandFinalize,
		MediaID: init.MediaID,
	}

	apiError := new(APIError)
	mediaFinalize := new(MediaFinalize)

	resp, err := m.sling.New().Base(twitterUploadAPI).Path("media/").Post("upload.json").BodyForm(params).Receive(mediaFinalize, apiError)

	return mediaFinalize, resp, relevantError(err, *apiError)

}

// UploadStatus Get the status from the uploaded asset
// https://developer.twitter.com/en/docs/media/upload-media/api-reference/get-media-upload-status
func (m *MediaService) UploadStatus(mediaID int64) (*MediaStatus, *http.Response, error) {
	apiError := new(APIError)
	mediaStatus := new(MediaStatus)

	resp, err := m.sling.New().Base(twitterUploadAPI).Path("media/").Get("upload.json").QueryStruct(MediaStatusParam{Command: twitterCommandStatus, MediaID: mediaID}).Receive(mediaStatus, apiError)

	return mediaStatus, resp, relevantError(err, *apiError)
}
