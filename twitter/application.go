package twitter

import (
	"github.com/dghubble/sling"
	"net/http"
	"strings"
	"time"
)

// ApplicationService provides a method for application rate limiting status.
type ApplicationService struct {
	sling *sling.Sling
}

// newApplicationService returns a new ApplicationService.
func newApplicationService(sling *sling.Sling) *ApplicationService {
	return &ApplicationService{
		sling: sling.Path("application/"),
	}
}

// RateLimitStatus wraps the rate limit status from Twitter.
type RateLimitStatus struct {
	RateLimitContext *RateLimitContext   `json:"rate_limit_context"`
	Resources        *RateLimitResources `json:"resources"`
}

// RateLimitContext provides the context deduced by the credentials.
type RateLimitContext struct {
	AccessToken string `json:"access_token"`
}

// RateLimitResources wraps the rate limit statuses per topic.
type RateLimitResources struct {
	Users    *UsersRateLimitResource    `json:"users"`
	Search   *SearchRateLimitResource   `json:"search"`
	Statuses *StatusesRateLimitResource `json:"statuses"`
	Help     *HelpRateLimitResource     `json:"help"`
}

// UsersRateLimitResource wraps the rate limit information for Users API endpoints.
type UsersRateLimitResource struct {
	ProfileBanner     *RateLimitResource `json:"/users/profile_banner"`
	SuggestionMembers *RateLimitResource `json:"/users/suggestions/:slug/members"`
	UserShow          *RateLimitResource `json:"/users/show/:id"`
	Suggestions       *RateLimitResource `json:"/users/suggestions"`
	Lookup            *RateLimitResource `json:"/users/lookup"`
	Search            *RateLimitResource `json:"/users/search"`
	Contributors      *RateLimitResource `json:"/users/contributors"`
	Contributees      *RateLimitResource `json:"/users/contributees"`
	SuggestionsSlug   *RateLimitResource `json:"/users/suggestions/:slug"`
}

// SearchRateLimitResource wraps the rate limit information for Search API endpoints.
type SearchRateLimitResource struct {
	Tweets *RateLimitResource `json:"/search/tweets"`
}

// StatusesRateLimitResource wraps the rate limit information for Statuses API endpoints.
type StatusesRateLimitResource struct {
	MentionsTimeline *RateLimitResource `json:"/statuses/mentions_timeline"`
	Lookup           *RateLimitResource `json:"/statuses/lookup"`
	Show             *RateLimitResource `json:"/statuses/show/:id"`
	Oembed           *RateLimitResource `json:"/statuses/oembed"`
	RetweetersIds    *RateLimitResource `json:"/statuses/retweeters/ids"`
	HomeTimeline     *RateLimitResource `json:"/statuses/home_timeline"`
	UserTimeline     *RateLimitResource `json:"/statuses/user_timeline"`
	Retweets         *RateLimitResource `json:"/statuses/retweets/:id"`
	RetweetsOfMe     *RateLimitResource `json:"/statuses/retweets_of_me"`
}

// HelpRateLimitResource wraps the rate limit information for Help API endpoints.
type HelpRateLimitResource struct {
	Privacy       *RateLimitResource `json:"/help/privacy"`
	ToS           *RateLimitResource `json:"/help/tos"`
	Configuration *RateLimitResource `json:"/help/configuration"`
	Languages     *RateLimitResource `json:"/help/languages"`
}

// RateLimitResource wraps a single rate limit status.
type RateLimitResource struct {
	Limit     int   `json:"limit"`
	Remaining int   `json:"remaining"`
	Reset     int64 `json:"reset"`
}

// ResetTimestamp provides the rate limiting reset time.
func (resource *RateLimitResource) ResetTimestamp() time.Time {
	return time.Unix(resource.Reset, 0)
}

type rateLimitStatusParams struct {
	Resources string `url:"resources,omitempty"`
}

// RateLimitStatus provides the rate limit status from Twitter.
func (s *ApplicationService) RateLimitStatus() (*RateLimitStatus, *http.Response, error) {
	params := rateLimitStatusParams{
		Resources: strings.Join([]string{"help", "users", "search", "statuses"}, ","),
	}
	rateLimitStatus := new(RateLimitStatus)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("rate_limit_status.json").QueryStruct(params).Receive(rateLimitStatus, apiError)
	return rateLimitStatus, resp, relevantError(err, *apiError)
}
