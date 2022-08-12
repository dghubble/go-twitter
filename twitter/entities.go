package twitter

// Entities represent metadata and context info parsed from Twitter components.
// https://developer.twitter.com/en/docs/tweets/data-dictionary/overview/entities-object
type Entities struct {
	Hashtags     []HashtagEntity `json:"hashtags"`
	Media        []MediaEntity   `json:"media"`
	Urls         []URLEntity     `json:"urls"`
	UserMentions []MentionEntity `json:"user_mentions"`
	Symbols      []SymbolEntity  `json:"symbols"`
	Polls        []PollEntity    `json:"polls"`
}

// HashtagEntity represents a hashtag which has been parsed from text.
type HashtagEntity struct {
	Indices Indices `json:"indices"`
	Text    string  `json:"text"`
}

// URLEntity represents a URL which has been parsed from text.
type URLEntity struct {
	URL         string     `json:"url"`
	DisplayURL  string     `json:"display_url"`
	ExpandedURL string     `json:"expanded_url"`
	Unwound     UnwoundURL `string:"unwound"`
	Indices     Indices    `json:"indices"`
}

// UnwoundURL represents an enhanced URL
// https://developer.twitter.com/en/docs/twitter-api/enterprise/enrichments/overview/expanded-and-enhanced-urls
type UnwoundURL struct {
	URL         string `json:"url"`
	Status      int    `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// MediaEntity represents media elements associated with a Tweet.
type MediaEntity struct {
	URLEntity
	ID                int64      `json:"id"`
	IDStr             string     `json:"id_str"`
	MediaURL          string     `json:"media_url"`
	MediaURLHttps     string     `json:"media_url_https"`
	SourceStatusID    int64      `json:"source_status_id"`
	SourceStatusIDStr string     `json:"source_status_id_str"`
	Type              string     `json:"type"`
	Sizes             MediaSizes `json:"sizes"`
	VideoInfo         VideoInfo  `json:"video_info"`
}

// MentionEntity represents Twitter user mentions parsed from text.
type MentionEntity struct {
	Indices    Indices `json:"indices"`
	ID         int64   `json:"id"`
	IDStr      string  `json:"id_str"`
	Name       string  `json:"name"`
	ScreenName string  `json:"screen_name"`
}

// SymbolEntity represents a symbol (e.g. $twtr) which has been parsed from text.
// https://developer.twitter.com/en/docs/twitter-api/v1/data-dictionary/object-model/entities#symbols
type SymbolEntity struct {
	Indices Indices `json:"indices"`
	Text    string  `json:"text"`
}

// PollEntity represents a Twitter Poll from a Tweet.
// Note that poll metadata is only available with enterprise.
// https://developer.twitter.com/en/docs/twitter-api/v1/data-dictionary/object-model/entities#polls
type PollEntity struct {
	Options         []PollOption `json:"options"`
	EndDateTime     string       `json:"end_datetime"`
	DurationMinutes string       `json:"duration_minutes"`
}

// PollOption represents a position option in a PollEntity.
type PollOption struct {
	Position int    `json:"position"`
	Text     string `json:"text"`
}

// UserEntities contain Entities parsed from User url and description fields.
// https://developer.twitter.com/en/docs/tweets/data-dictionary/overview/entities-object#mentions
type UserEntities struct {
	URL         Entities `json:"url"`
	Description Entities `json:"description"`
}

// ExtendedEntity contains media information.
// https://developer.twitter.com/en/docs/tweets/data-dictionary/overview/extended-entities-object
type ExtendedEntity struct {
	Media []MediaEntity `json:"media"`
}

// Indices represent the start and end offsets within text.
type Indices [2]int

// Start returns the index at which an entity starts, inclusive.
func (i Indices) Start() int {
	return i[0]
}

// End returns the index at which an entity ends, exclusive.
func (i Indices) End() int {
	return i[1]
}

// MediaSizes contain the different size media that are available.
// https://developer.twitter.com/en/docs/tweets/data-dictionary/overview/entities-object#media-size
type MediaSizes struct {
	Thumb  MediaSize `json:"thumb"`
	Large  MediaSize `json:"large"`
	Medium MediaSize `json:"medium"`
	Small  MediaSize `json:"small"`
}

// MediaSize describes the height, width, and resizing method used.
type MediaSize struct {
	Width  int    `json:"w"`
	Height int    `json:"h"`
	Resize string `json:"resize"`
}

// VideoInfo is available on video media objects.
type VideoInfo struct {
	AspectRatio    [2]int         `json:"aspect_ratio"`
	DurationMillis int            `json:"duration_millis"`
	Variants       []VideoVariant `json:"variants"`
}

// VideoVariant describes one of the available video formats.
type VideoVariant struct {
	ContentType string `json:"content_type"`
	Bitrate     int    `json:"bitrate"`
	URL         string `json:"url"`
}
