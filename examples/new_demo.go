package main

import (
	"encoding/csv"
	"fmt"
	"os"
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

// Add a method to the user
func (m *Message) CSV() []string {
	return []string{strconv.Itoa(m.CreatedAt), m.Text}
}

// Add a method to the user
func (u *User) CSV() [][]string {
	dms := make([][]string, 0)
	for see := range u.Messages {
		dms = append(dms, u.Messages[see].CSV())
	}
	return dms
}

//userParser to collate user messages
func userParser(dms *twitter.DirectMessageEvent, client *twitter.Client) bool {
	userIDs := []int64{}
	for m := range dms.Events {
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
	for k := range users {
		// fmt.Println(len(users[k].Messages))
		i, _ := strconv.Atoi(k)
		userIDs = append(userIDs, int64(i))
	}
	fmt.Println(userIDs)
	fmt.Println("Looking up associated user data")
	usersLookup, _, _ := client.Users.Lookup(&twitter.UserLookupParams{UserID: userIDs})
	for user := range usersLookup {
		u := usersLookup[user]
		uID := u.IDStr
		uu := users[uID]
		uu.ScreenName = u.ScreenName
		users[uID] = uu
	}
	for u := range users {
		user := users[u]
		sort.Slice(user.Messages, func(eye, jay int) bool {
			return user.Messages[eye].CreatedAt < user.Messages[jay].CreatedAt
		})
	}
	fmt.Println("Writing user data to CSV")
	for yew := range users {
		user := users[yew]
		file, _ := os.OpenFile(fmt.Sprintf("/tmp/twitter/%s", user.ScreenName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
		writer := csv.NewWriter(file)
		defer writer.Flush()
		writer.WriteAll(user.CSV())
	}
	users = make(map[string]User)
	userIDs = []int64{}
	return true
}

func main() {
	twitterEventsRestRateLimit := rate.New(1, time.Minute)
	twitterAPIDontKnowConsistencyRate := rate.New(1, time.Second*2)
	var twitterConfig TwitterConfig
	envconfig.Process("twitter", &twitterConfig)
	config := oauth1.NewConfig(twitterConfig.ConsumerKey, twitterConfig.ConsumerSecret)
	token := oauth1.NewToken(twitterConfig.AccessToken, twitterConfig.AccessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)
	// Twitter client
	client := twitter.NewClient(httpClient)
	// Initial Rest calls to the twitter events
	dms, _, _ := client.DirectMessages.GetEvents(&twitter.DirectMessageEventsGetParams{Count: 50})
	userParser(&dms, client)
	fmt.Println("Ran GetEvents at " + fmt.Sprintf(time.Now().String()))

	for i := 0; i <= 40; i++ {
		cursor := dms.NextCursor
		twitterEventsRestRateLimit.Wait()
		twitterAPIDontKnowConsistencyRate.Wait()
		dms, _, _ = client.DirectMessages.GetEvents(&twitter.DirectMessageEventsGetParams{NextCursor: dms.NextCursor, Count: 50})
		if len(dms) < 1 {
			twitterEventsRestRateLimit.Wait()
			dms, _, _ = client.DirectMessages.GetEvents(&twitter.DirectMessageEventsGetParams{NextCursor: cursor, Count: 50})
		}

		fmt.Println("Ran GetEvents at " + fmt.Sprintf(time.Now().String()))

		userParser(&dms, client)
	}
}
