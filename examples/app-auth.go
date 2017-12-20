package main

import (
	"fmt"
	"sort"
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
	SenderID   string
	ScreenName string
	Messages   []Message
}

// Message contains the raw DM content
type Message struct {
	Text      string
	CreatedAt int
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
		createdTimestamp, _ := strconv.Atoi(dms.Events[m].CreatedAt)
		user.Messages = append(user.Messages, Message{
			Text:      dms.Events[m].Message.Data.Text,
			CreatedAt: createdTimestamp,
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
	// initialRestCall := time.Now()

	// Initial Rest calls to the twitter events
	dms, _, _ := client.DirectMessages.GetEvents(&twitter.DirectMessageEventsGetParams{Count: 50})
	userParser(&dms)

	// fmt.Println(initialRestCall)

	for i := 1; i <= 1; i++ {
		twitterEventsRestRateLimit.Wait()
		dms, _, _ = client.DirectMessages.GetEvents(&twitter.DirectMessageEventsGetParams{NextCursor: dms.NextCursor, Count: 50})
		userParser(&dms)
	}

	userIDs := []int64{}
	for k := range users {
		// fmt.Println(len(users[k].Messages))
		i, _ := strconv.Atoi(k)
		userIDs = append(userIDs, int64(i))
	}

	usersLookup, _, _ := client.Users.Lookup(&twitter.UserLookupParams{UserID: userIDs})

	for user := range usersLookup {
		u := usersLookup[user]
		uID := u.IDStr
		uu := users[uID]
		uu.ScreenName = u.ScreenName
		users[uID] = uu
		// fmt.Println(u.ScreenName)
		// fmt.Println(uu)
	}

	fmt.Println(userIDs)
	// myUser := users["14259060"]

	for u := range users {
		user := users[u]
		sort.Slice(user.Messages, func(eye, jay int) bool {
			return user.Messages[eye].CreatedAt < user.Messages[jay].CreatedAt
		})
	}

	// fmt.Println(body)
	// fmt.Println(userIDs)
	// for m := range dms.Events {

	// 	dm := dms.Events[m]
	// 	fmt.Println(dm.Message.Data.Text)
	// 	fmt.Println(dm.Message.SenderID)
	// }
}
