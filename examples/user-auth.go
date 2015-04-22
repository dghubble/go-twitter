package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"os"
)

// Main makes User Auth (OAuth1) requests as a consumer on behalf of a user.
func main() {
	// read credentials from environment variables
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessTokenSecret == "" {
		panic("Missing required environment variable")
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(token)

	// twitter client
	client := twitter.NewClient(httpClient)

	// Home Timeline
	homeTimelineParams := &twitter.HomeTimelineParams{Count: 2}
	tweets, _, _ := client.Timelines.HomeTimeline(homeTimelineParams)
	fmt.Printf("User's HOME TIMELINE:\n%+v\n", tweets)

	// Mention Timeline
	mentionTimelineParams := &twitter.MentionTimelineParams{Count: 2}
	tweets, _, _ = client.Timelines.MentionTimeline(mentionTimelineParams)
	fmt.Printf("User's MENTION TIMELINE:\n%+v\n", tweets)

	// Retweets of Me Timeline
	retweetTimelineParams := &twitter.RetweetsOfMeTimelineParams{Count: 2}
	tweets, _, _ = client.Timelines.RetweetsOfMeTimeline(retweetTimelineParams)
	fmt.Printf("User's 'RETWEETS OF ME' TIMELINE:\n%+v\n", tweets)

	// Update (POST!) Tweet (uncomment to run)
	// tweet, _, _ := client.Statuses.Update("just setting up my twttr", nil)
	// fmt.Printf("Posted Tweet\n%v\n", tweet)
}
