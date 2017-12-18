package main

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/kelseyhightower/envconfig"
)

// TwitterConfig from ENV Vars
type TwitterConfig struct {
	AccessToken    *string `envconfig:"access_token"`
	AccessSecret   *string `envconfig:"access_secret"`
	ConsumerKey    *string `envconfig:"consumer_key"`
	ConsumerSecret *string `envconfig:"consumer_secret"`
}

// User to container message data and user information
type User struct {
	SenderID string
	Messages []Message
}

// Message contains the raw DM content
type Message struct {
	Text      string
	CreatedAt string
}

func main() {
	// twitterEventsRestRateLimit := rate.New(15, time.Minute)

	var twitterConfig TwitterConfig
	envconfig.Process("twitter", &twitterConfig)
	config := oauth1.NewConfig(*twitterConfig.ConsumerKey, *twitterConfig.ConsumerSecret)
	token := oauth1.NewToken(*twitterConfig.AccessToken, *twitterConfig.AccessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)
	// Twitter client
	client := twitter.NewClient(httpClient)
	// When did we first start our rest calls
	// initialRestCall := time.Now()

	users := make(map[string]User)

	// Initial Rest calls to the twitter events
	dms, _, _ := client.DirectMessages.GetEvents(&twitter.DirectMessageEventsGetParams{Count: 20})

	for m := range dms.Events {
		fmt.Printf("%s - %s\n", dms.Events[m].Message.SenderID, dms.Events[m].Message.Data.Text)
		if _, ok := users[dms.Events[m].Message.SenderID]; !ok {
			// User Not found, create one
			user := User{
				Messages: make([]Message, 0),
			}
			users[dms.Events[m].Message.SenderID] = user
		}
		user := users[dms.Events[m].Message.SenderID]
		user.Messages = append(user.Messages, Message{
			Text:      dms.Events[m].Message.Data.Text,
			CreatedAt: dms.Events[m].CreatedAt,
		})
		users[dms.Events[m].Message.SenderID] = user
	}

	fmt.Println(users)
	// fmt.Println(dms)
	// fmt.Println(dms.NextCursor)

	// for i := 1; i <= 1; i++ {
	// 	twitterEventsRestRateLimit.Wait()
	// 	dms, _, _ := client.DirectMessages.GetEvents(&twitter.DirectMessageEventsGetParams{NextCursor: dms.NextCursor, Count: 1})
	// 	fmt.Println(dms.NextCursor)
	// }

	// for m := range dms.Events {

	// 	dm := dms.Events[m]
	// 	fmt.Println(dm.Message.Data.Text)
	// 	fmt.Println(dm.Message.SenderID)
	// }
}
