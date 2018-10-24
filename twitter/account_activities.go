package twitter

import (
	"github.com/dghubble/sling"
	"net/http"
)

// AccountActivityService provides methods for accessing Twitter's account
// activities endpoints
type AccountActivityService struct {
	sling *sling.Sling
}

// AccountActivityRegisterWebhookParams are the parameters used for registering
// a webhook on the account activities API.
type AccountActivityRegisterWebhookParams struct {
	EnvName string
	URL     string `url:"url"`
}

// CreateSubscriptionParams are the parameters used for subscribing to events
// for a given user
type AccountActivityCreateSubscriptionParams struct {
	EnvName string
}

// AccountActivityWebhook contains information about a webhook created on the account activity
// API
type AccountActivityWebhook struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	Valid     bool   `json:"valid"`
	CreatedAt string `json:"created_at"`
}

func newAccountActivityService(sling *sling.Sling) *AccountActivityService {
	return &AccountActivityService{
		sling: sling.Path("account_activity/"),
	}
}

// RegisterWebhook registers a given URL as a webhook for the account
// activities API
func (s *AccountActivityService) RegisterWebhook(params *AccountActivityRegisterWebhookParams) (*AccountActivityWebhook, *http.Response, error) {
	if params == nil {
		params = &AccountActivityRegisterWebhookParams{}
	}
	apiError := new(APIError)
	webhook := new(AccountActivityWebhook)
	resp, err := s.sling.New().Post("all/"+params.EnvName+"/webhooks.json").BodyForm(params).Receive(webhook, apiError)

	return webhook, resp, relevantError(err, *apiError)
}

// CreateSubscription subscribes the application's webhooks to all of the
// authenticated user's events
func (s *AccountActivityService) CreateSubscription(params *AccountActivityCreateSubscriptionParams) (*http.Response, error) {
	if params == nil {
		params = &AccountActivityCreateSubscriptionParams{}
	}
	apiError := new(APIError)
	resp, err := s.sling.New().Post("all/"+params.EnvName+"/subscriptions.json").BodyForm(params).Receive(nil, apiError)

	return resp, relevantError(err, *apiError)
}
