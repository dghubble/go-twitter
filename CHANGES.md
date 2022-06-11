# go-twitter Changelog

Notable changes over time. Note, `go-twitter` does not follow a semver release cycle since it may change whenever the Twitter API changes (external).

## 06/2022

* Forked off from [original](https://github.com/dghubble/go-twitter) so we
can start taking pull requests.

## 07/2019

* Add Go module support ([#143](https://github.com/dghubble/go-twitter/pull/143))

## 11/2018

* Add `DirectMessageService` support for the new Twitter Direct Message Events [API](https://developer.twitter.com/en/docs/direct-messages/api-features) ([#125](https://github.com/dghubble/go-twitter/pull/125))
  * Add `EventsNew` method for sending a Direct Message event
  * Add `EventsShow` method for getting a single Direct Message event
  * Add `EventsList` method for listing recent Direct Message events
  * Add`EventsDestroy` method for destroying a Direct Message event

