

# go-twitter [![Build Status](https://travis-ci.org/dghubble/go-twitter.png)](https://travis-ci.org/dghubble/go-twitter) [![Coverage](http://gocover.io/_badge/github.com/dghubble/go-twitter/twitter)](http://gocover.io/github.com/dghubble/go-twitter/twitter) [![GoDoc](http://godoc.org/github.com/dghubble/go-twitter?status.png)](http://godoc.org/github.com/dghubble/go-twitter)
<img align="right" src="http://storage.googleapis.com/dghubble/gopher-on-bird.png">

go-twitter is a Go client library for the [Twitter API](https://dev.twitter.com/rest/public).

### Features

* Twitter API services:
    * StatusService
    * TimelineService
    * UserService
    * FollowerService

## Install

    go get github.com/dghubble/go-twitter/twitter

## Documentation

Read [GoDoc](https://godoc.org/github.com/dghubble/go-twitter/twitter)

## Usage

The `twitter` package provides a `Client` for accessing the Twitter API. Here are some example requests.

```go
// Twitter client
client := twitter.NewClient(httpClient)

// Home Timeline
tweets, resp, err := client.Timelines.HomeTimeline(&HomeTimelineParams{})

// Send a Tweet
tweet, resp, err := client.Statuses.Update("just setting up my twttr", nil)

// Status Show
tweet, resp, err := client.Statuses.Show(585613041028431872, nil)

// User Show
params := &twitter.UserShowParams{ScreenName: "dghubble"}
user, resp, err := client.Users.Show(params)

// Followers
followers, resp, err := client.Followers.List(&FollowerListParams{})
```

Required parameters are passed as positional arguments. Optional parameters are passed in a typed params struct (or pass nil).

## Authentication

By design, the Twitter Client accepts any `http.Client` so user auth (OAuth1) or application auth (OAuth2) requests can be made by using the appropriate authenticated client. Use the [dghubble/oauth1](https://github.com/dghubble/oauth1) and [golang/oauth2](https://github.com/golang/oauth2/) packages to obtain an `http.Client` which transparently authorizes requests.

For example, make requests as a consumer application on behalf of a user who has granted access, with OAuth1.

```go
// OAuth1
import (
    "github.com/dghubble/go-twitter/twitter"
    "github.com/dghubble/oauth1"
)

config := oauth1.NewConfig("consumerKey", "consumerSecret")
token := oauth1.NewToken("accessToken", "accessSecret")
// http.Client will automatically authorize Requests
httpClient := config.Client(oauth1.NoContext, token)

// twitter client
client := twitter.NewClient(httpClient)
```

If no user auth context is needed, make requests as your application with application auth.

```go
// OAuth2
import (
    "github.com/dghubble/go-twitter/twitter"
    "golang.org/x/oauth2"
)

config := &oauth2.Config{}
token := &oauth2.Token{AccessToken: accessToken}
// http.Client will automatically authorize Requests
httpClient := config.Client(oauth2.NoContext, token)

// twitter client
client := twitter.NewClient(httpClient)
```

To implement Login with Twitter for web or mobile, see the gologin [package](https://github.com/dghubble/gologin) and [examples](https://github.com/dghubble/gologin/tree/master/examples/twitter).

## Contributing

See the [Contributing Guide](https://gist.github.com/dghubble/be682c123727f70bcfe7).

## License

[MIT License](LICENSE)
