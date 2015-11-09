package twitter

// DirectMessage is a direct message to a single recipient.
type DirectMessage struct {
	CreatedAt           string    `json:"created_at"`
	Entities            *Entities `json:"entities"`
	ID                  int64     `json:"id"`
	IDStr               string    `json:"id_str"`
	Recipient           *User     `json:"recipient"`
	RecipientID         int64     `json:"recipient_id"`
	RecipientScreenName string    `json:"recipient_screen_name"`
	Sender              *User     `json:"sender"`
	SenderID            int64     `json:"sender_id"`
	SenderScreenName    string    `json:"sender_screen_name"`
	Text                string    `json:"text"`
}
