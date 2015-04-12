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
	// OAuth 2 Bearer Access Token
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	if accessToken == "" {
		panic("Set the TWITTER_ACCESS_TOKEN environment variable")
	}
	ts := NewTokenSource(&oauth2.Token{AccessToken: accessToken})
	appAuthClient := oauth2.NewClient(oauth2.NoContext, ts)

	// Twitter

	client := twitter.NewClient(appAuthClient)

	// user show
	params := &twitter.UserShowParams{ScreenName: "dghubble"}
	user, resp, err := client.Users.Show(params)
	fmt.Printf("%+v\n", user)
	fmt.Println(resp, err)
}

// golang/oauth2

type tokenSource struct {
	token *oauth2.Token
}

func NewTokenSource(token *oauth2.Token) *tokenSource {
	return &tokenSource{
		token: token,
	}
}

func (t *tokenSource) Token() (*oauth2.Token, error) {
	return t.token, nil
}
