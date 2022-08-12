module github.com/dghubble/go-twitter/examples

go 1.17

require (
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f
	github.com/dghubble/go-twitter v0.0.0-20220716034336-7f63262ef83a
	github.com/dghubble/oauth1 v0.6.0
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
)

require (
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/dghubble/sling v1.4.0 // indirect
	github.com/golang/protobuf v1.2.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	golang.org/x/net v0.0.0-20190108225652-1e06a53dbb7e // indirect
	google.golang.org/appengine v1.4.0 // indirect
)

replace github.com/dghubble/go-twitter/twitter => ../
