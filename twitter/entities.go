package twitter

// https://dev.twitter.com/overview/api/entities
// TODO: symbols, extended_entities
type Entities struct {
	Hashtags     []HashtagEntity `json:"hashtags"`
	Media        []MediaEntity   `json:"media"`
	Urls         []UrlEntity     `json:"urls"`
	UserMentions []MentionEntity `json:"user_mentions"`
}

type HashtagEntity struct {
	Indices Indices `json:"indices"`
	Text    string  `json:"text"`
}

type UrlEntity struct {
	Indices     Indices `json:"indices"`
	DisplayUrl  string  `json:"display_url"`
	ExpandedUrl string  `json:"expanded_url"`
	Url         string  `json:"url"`
}

// TODO: add Sizes
type MediaEntity struct {
	UrlEntity
	Id                int64  `json:"id"`
	IdStr             string `json:"id_str"`
	MediaUrl          string `json:"media_url"`
	MediaUrlHttps     string `json:"media_url_https"`
	SourceStatusId    int64  `json:"source_status_id"`
	SourceStatusIdStr string `json:"source_status_id_str"`
	Type              string `json:"type"`
}

type MentionEntity struct {
	Indices    Indices `json:"indices"`
	Id         int64   `json:"id"`
	IdStr      string  `json:"id_str"`
	Name       string  `json:"name"`
	ScreenName string  `json:"screen_name"`
}

// UserEntities contain Entities parsed from User url and description fields.
// https://dev.twitter.com/overview/api/entities-in-twitter-objects#users
type UserEntities struct {
	Url         Entities `json:"url"`
	Description Entities `json:"description"`
}

// Indices represent the start and end offsets within text.
type Indices [2]int

func (i Indices) Start() int {
	return i[0]
}

func (i Indices) End() int {
	return i[1]
}
