

# go-twitter [![Build Status](https://travis-ci.org/dghubble/go-twitter.png)](https://travis-ci.org/dghubble/go-twitter) [![GoDoc](http://godoc.org/github.com/dghubble/go-twitter?status.png)](http://godoc.org/github.com/dghubble/go-twitter)

go-twitter is an (IN PROGRESS) Go client library for the [Twitter API](https://dev.twitter.com/rest/public).

## Install

    go get github.com/dghubble/go-twitter/twitter

## Documentation

Read [GoDoc](https://godoc.org/github.com/dghubble/go-twitter/twitter)

## Usage

The `twitter` package contains the Twitter API services which can be accessed through the client.

```go
// twitter client
client := twitter.NewClient(authClient)

// Home Timeline
tweets, resp, err := client.Timelines.HomeTimeline(nil)

// User Show
params := &twitter.UserShowParams{ScreenName: "dghubble"}
user, resp, err := client.Users.Show(params)

// Status Show (Tweet show)
tweet, resp, err := client.Statuses.Show(585613041028431872, nil)
```

## Authentication

The `twitter` package does not directly handle OAuth authentication. For OAuth1 "user-auth" requests, the [dghubble/oauth1](https://github.com/dghubble/oauth1) implementation is recommended. For OAuth2 "app-auth" requests, use the official [golang/oauth2](https://github.com/golang/oauth2/) library.

To make requests as a consumer on behalf of a user who has granted access, "user-auth" (OAuth1) should be used. Create an `http.Client` which will sign requests on behalf of a user:

```go
// OAuth1 (user-auth)
import "github.com/dghubble/oauth1"

config := oauth1.NewConfig(consumerKey, consumerSecret)
token := oauth1.NewToken(accessToken, accessTokenSecret)
// OAuth1 http.Client will automatically authorize Requests
httpClient := config.Client(token)

// twitter client
client := twitter.NewClient(authClient)
```

Alternately, if you do not need to make requests with user context, Twitter "app-auth" (OAuth2) allows requests to be made as an application. Create an `http.Client` which will make OAuth2 requests.

```go
// OAuth2 (app-auth)
import "golang.org/x/oauth2"

type tokenSource struct {
    token *oauth2.Token
}

func (t *tokenSource) Token() (*oauth2.Token, error) {
    return t.token, nil
}

ts := &tokenSource{&oauth2.Token{AccessToken: accessToken}}
// OAuth2 http.Client will automatically authorize Requests
httpClient := oauth2.NewClient(oauth2.NoContext, ts)

// twitter client
client := twitter.NewClient(authClient)
```

Now use your client to make requests to Twitter!

## API Services

The `twitter` Client currently provides the following API services:

* StatusService
* TimelineService
* UserService
