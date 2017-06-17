package twitter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStopped(t *testing.T) {
	done := make(chan struct{})
	assert.False(t, stopped(done))
	close(done)
	assert.True(t, stopped(done))
}

func TestSleepOrDone_Sleep(t *testing.T) {
	wait := time.Nanosecond * 20
	done := make(chan struct{})
	completed := make(chan struct{})
	go func() {
		sleepOrDone(wait, done)
		close(completed)
	}()
	// wait for goroutine SleepOrDone to sleep
	assertDone(t, completed, defaultTestTimeout)
}

func TestSleepOrDone_Done(t *testing.T) {
	wait := time.Second * 5
	done := make(chan struct{})
	completed := make(chan struct{})
	go func() {
		sleepOrDone(wait, done)
		close(completed)
	}()
	// close done, interrupting SleepOrDone
	close(done)
	// assert that SleepOrDone exited, closing completed
	assertDone(t, completed, defaultTestTimeout)
}
