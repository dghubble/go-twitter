package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/beefsack/go-rate"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/kelseyhightower/envconfig"
)

// TwitterConfig from ENV Vars
type TwitterConfig struct {
	AccessToken    string `envconfig:"access_token"`
	AccessSecret   string `envconfig:"access_secret"`
	ConsumerKey    string `envconfig:"consumer_key"`
	ConsumerSecret string `envconfig:"consumer_secret"`
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

var users = make(map[string]User)

//userParser to collate user messages
func userParser(dms *twitter.DirectMessageEvent) bool {
	for m := range dms.Events {
		// fmt.Printf("%s - %s\n", dms.Events[m].Message.SenderID, dms.Events[m].Message.Data.Text)
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
	return true
}

func main() {
	twitterEventsRestRateLimit := rate.New(1, time.Second)

	var twitterConfig TwitterConfig
	envconfig.Process("twitter", &twitterConfig)
	config := oauth1.NewConfig(twitterConfig.ConsumerKey, twitterConfig.ConsumerSecret)
	token := oauth1.NewToken(twitterConfig.AccessToken, twitterConfig.AccessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)
	// Twitter client
	client := twitter.NewClient(httpClient)
	// When did we first start our rest calls
	initialRestCall := time.Now()

	// Initial Rest calls to the twitter events
	dms, _, _ := client.DirectMessages.GetEvents(&twitter.DirectMessageEventsGetParams{Count: 50})
	userParser(&dms)

	fmt.Println(initialRestCall)

	for i := 1; i <= 3; i++ {
		twitterEventsRestRateLimit.Wait()
		dms, _, _ = client.DirectMessages.GetEvents(&twitter.DirectMessageEventsGetParams{NextCursor: dms.NextCursor, Count: 50})
		fmt.Println(time.Now())
		fmt.Println(dms.NextCursor)
		userParser(&dms)
	}

	userIDs := []int64{}
	for k := range users {
		fmt.Println(len(users[k].Messages))
		i, _ := strconv.Atoi(k)
		userIDs = append(userIDs, int64(i))
	}

	// userLookup, _, _ := client.Users.Lookup(&twitter.UserLookupParams{UserID: userIDs})
	fmt.Println(userIDs)
	// for m := range dms.Events {

	// 	dm := dms.Events[m]
	// 	fmt.Println(dm.Message.Data.Text)
	// 	fmt.Println(dm.Message.SenderID)
	// }
}
