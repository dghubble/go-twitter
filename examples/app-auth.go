// Application authentication example
// Run with: go run examples/app-auth.go --api-key=XXX --api-secret=XXX

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	flags := flag.NewFlagSet("app-auth", flag.ExitOnError)
	apiKey := flags.String("api-key", "", "Twitter API Key")
	apiSecret := flags.String("api-secret", "", "Twitter API Secret")

	flags.Parse(os.Args[1:])
	flagutil.SetFlagsFromEnv(flags, "TWITTER")

	if *apiKey == "" {
		log.Fatal("Twitter API key (--api-key) required")
	}

	if *apiSecret == "" {
		log.Fatal("Twitter API secret (--api-secret) required")
	}

	// Build 2 legged oauth config
	config := &clientcredentials.Config{ClientID: *apiKey,
		ClientSecret: *apiSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token"}

	// OAuth2 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth2.NoContext)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// user show
	userShowParams := &twitter.UserShowParams{ScreenName: "golang"}
	user, _, _ := client.Users.Show(userShowParams)
	fmt.Printf("USERS SHOW:\n%+v\n", user)

	// users lookup
	userLookupParams := &twitter.UserLookupParams{ScreenName: []string{"golang", "gophercon"}}
	users, _, _ := client.Users.Lookup(userLookupParams)
	fmt.Printf("USERS LOOKUP:\n%+v\n", users)

	// status show
	statusShowParams := &twitter.StatusShowParams{}
	tweet, _, _ := client.Statuses.Show(584077528026849280, statusShowParams)
	fmt.Printf("STATUSES SHOW:\n%+v\n", tweet)

	// statuses lookup
	statusLookupParams := &twitter.StatusLookupParams{ID: []int64{20}}
	tweets, _, _ := client.Statuses.Lookup([]int64{573893817000140800}, statusLookupParams)
	fmt.Printf("STATUSES LOOKUP:\n%+v\n", tweets)

	// oEmbed status
	statusOembedParams := &twitter.StatusOEmbedParams{ID: 691076766878691329, MaxWidth: 500}
	oembed, _, _ := client.Statuses.OEmbed(statusOembedParams)
	fmt.Printf("OEMBED TWEET:\n%+v\n", oembed)

	// user timeline
	userTimelineParams := &twitter.UserTimelineParams{ScreenName: "golang", Count: 2}
	tweets, _, _ = client.Timelines.UserTimeline(userTimelineParams)
	fmt.Printf("USER TIMELINE:\n%+v\n", tweets)

	// search tweets
	searchTweetParams := &twitter.SearchTweetParams{
		Query: "happy birthday",
		Count: 3,
	}
	search, _, _ := client.Search.Tweets(searchTweetParams)
	fmt.Printf("SEARCH TWEETS:\n%+v\n", search)
	//fmt.Printf("SEARCH METADATA:\n%+v\n", search.SearchMetadata)
}
