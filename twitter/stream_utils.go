package twitter

import (
	"time"
)

// stopped returns true if the done channel receives, false otherwise.
func stopped(done <-chan struct{}) bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

// sleepOrDone pauses the current goroutine until the done channel receives
// or until at least the duration d has elapsed, whichever comes first. This
// is similar to time.Sleep(d), except it can be interrupted.
func sleepOrDone(d time.Duration, done <-chan struct{}) {
	select {
	case <-time.After(d):
		return
	case <-done:
		return
	}
}
