
# go-twitter Examples

## User Auth (OAuth1)

A user (OAuth1) access token grants a consumer application access to a user's  Twitter resources. Set the consumer key and secret and the access token and secret as environment variables.

    export TWITTER_CONSUMER_KEY=xxx
    export TWITTER_CONSUMER_SECRET=xxx
    export TWITTER_ACCESS_TOKEN=xxx
    export TWITTER_ACCESS_TOKEN_SECRET=xxx

Make requests as the application, on behalf of the user by running:

    go run user-auth.go

to show the home timeline, mention timeline, and retweets timeline.

## App Auth (OAuth2)

An application "app-auth" (OAuth2) access token allows an application to make Twitter API requests for public content, with rate limits counting against the app itself. App-auth requests can be made to API endpoints which do not require a user context.

    export TWITTER_APP_ACCESS_TOKEN=xxx

Make requests as the application by running:

    go run app-auth.go

to load some public Users and Tweets.