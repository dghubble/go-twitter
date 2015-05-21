

# go-twitter [![Build Status](https://travis-ci.org/dghubble/go-twitter.png)](https://travis-ci.org/dghubble/go-twitter) [![Coverage](http://gocover.io/_badge/github.com/dghubble/go-twitter/twitter)](http://gocover.io/github.com/dghubble/go-twitter/twitter) [![GoDoc](http://godoc.org/github.com/dghubble/go-twitter?status.png)](http://godoc.org/github.com/dghubble/go-twitter)
<img align="right" src="http://storage.googleapis.com/dghubble/gopher-on-bird.png">

go-twitter is an (in progress) Go client library for the [Twitter API](https://dev.twitter.com/rest/public).

### Features

* Package `twitter` provides Twitter API services:
    * StatusService
    * TimelineService
    * UserService

## Install

    go get github.com/dghubble/go-twitter/twitter

## Documentation

Read [GoDoc](https://godoc.org/github.com/dghubble/go-twitter/twitter)

## Usage

The `twitter` package contains Twitter API services which can be accessed through the client.

```go
// twitter client
client := twitter.NewClient(authClient)

// Home Timeline
tweets, resp, err := client.Timelines.HomeTimeline(nil)

// Send a Tweet
tweet, resp, err := client.Statuses.Update("just setting up my twttr", nil)

// Status Show
tweet, resp, err := client.Statuses.Show(585613041028431872, nil)

// User Show
params := &twitter.UserShowParams{ScreenName: "dghubble"}
user, resp, err := client.Users.Show(params)
```

Required parameters are passed as positional arguments. Optional parameters are passed via a typed params struct for each endpoint.

Method names match the Twitter API endpoint names, except timeline-type endpoints are provided by `TimelineService` rather than `StatusService`.

## Authentication

By design, the `twitter` package client is decoupled from authentication concerns. Twitter "user auth" and "app auth" endpoints require [OAuth1](https://tools.ietf.org/html/rfc5849) and [OAuth2](https://tools.ietf.org/html/rfc6749), respectively. Use the [dghubble/oauth1](https://github.com/dghubble/oauth1) and [golang/oauth2](https://github.com/golang/oauth2/) libraries to obtain an `http.Client`, which transparently handles authorizing requests.

For example, make requests as a consumer on behalf of a user who has granted access, with OAuth1 "user auth":

```go
// OAuth1
import "github.com/dghubble/oauth1"

config := oauth1.NewConfig(consumerKey, consumerSecret)
token := oauth1.NewToken(accessToken, accessTokenSecret)
// OAuth1 http.Client will automatically authorize Requests
httpClient := config.Client(token)

// twitter client
client := twitter.NewClient(authClient)
```

If no user context is needed, make requests as your application with app-auth (OAuth2):

```go
// OAuth2
import "golang.org/x/oauth2"

config := &oauth2.Config{}
token := &oauth2.Token{AccessToken: accessToken}
// OAuth2 http.Client will automatically authorize Requests
httpClient := config.Client(oauth2.NoContext, token)

// twitter client
client := twitter.NewClient(authClient)
```

Now use your `twitter` client to make requests.

## License

[MIT License](LICENSE)
