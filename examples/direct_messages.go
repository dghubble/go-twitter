package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
	consumerKey := flags.String("consumer-key", "", "Twitter Consumer Key")
	consumerSecret := flags.String("consumer-secret", "", "Twitter Consumer Secret")
	accessToken := flags.String("access-token", "", "Twitter Access Token")
	accessSecret := flags.String("access-secret", "", "Twitter Access Secret")
	flags.Parse(os.Args[1:])
	flagutil.SetFlagsFromEnv(flags, "TWITTER")

	if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// List most recent 10 Direct Messages
	messages, _, err := client.DirectMessages.EventsList(
		&twitter.DirectMessageEventsListParams{Count: 10},
	)
	fmt.Println("User's DIRECT MESSAGES:")
	if err != nil {
		log.Fatal(err)
	}
	for _, event := range messages.Events {
		fmt.Printf("%+v\n", event)
		fmt.Printf("  %+v\n", event.Message)
		fmt.Printf("  %+v\n", event.Message.Data)
	}

	// Show Direct Message event
	event, _, err := client.DirectMessages.EventsShow("1066903366071017476", nil)
	fmt.Printf("DM Events Show:\n%+v, %v\n", event.Message.Data, err)

	// Create Direct Message event
	/*
		event, _, err = client.DirectMessages.EventsNew(&twitter.DirectMessageEventsNewParams{
			Event: &twitter.DirectMessageEvent{
				Type: "message_create",
				Message: &twitter.DirectMessageEventMessage{
					Target: &twitter.DirectMessageTarget{
						RecipientID: "2856535627",
					},
					Data: &twitter.DirectMessageData{
						Text: "testing",
					},
				},
			},
		})
		fmt.Printf("DM Event New:\n%+v, %v\n", event, err)
	*/

	// Destroy Direct Message event
	//_, err = client.DirectMessages.EventsDestroy("1066904217049133060")
	//fmt.Printf("DM Events Delete:\n err: %v\n", err)
}
