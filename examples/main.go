package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"os"
)

func main() {
	// create an http.Client which handles authentication

	// for Twitter "app-only auth" use golang/oauth2 http client
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	if accessToken == "" {
		panic("Missing TWITTER_ACCESS_TOKEN environment variable")
	}
	ts := &tokenSource{&oauth2.Token{AccessToken: accessToken}}
	appAuthClient := oauth2.NewClient(oauth2.NoContext, ts)

	// Twitter

	client := twitter.NewClient(appAuthClient)

	// user show
	userShowParams := &twitter.UserShowParams{ScreenName: "dghubble"}
	user, _, _ := client.Users.Show(userShowParams)
	fmt.Printf("user/show:\n%+v\n", user)

	// status show
	statusShowParams := &twitter.StatusShowParams{}
	tweet, _, _ := client.Statuses.Show(584077528026849280, statusShowParams)
	fmt.Printf("statuses/show:\n%+v\n", tweet)

	// user timeline
	userTimelineParams := &twitter.UserTimelineParams{ScreenName: "golang"}
	tweets, _, _ := client.Timelines.UserTimeline(userTimelineParams)
	fmt.Printf("statuses/user_timeline:\n%+v\n", tweets)
}

// golang/oauth2

type tokenSource struct {
	token *oauth2.Token
}

func (t *tokenSource) Token() (*oauth2.Token, error) {
	return t.token, nil
}
