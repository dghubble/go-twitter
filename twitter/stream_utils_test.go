package twitter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestScanLines(t *testing.T) {
	cases := []struct {
		input   []byte
		atEOF   bool
		advance int
		token   []byte
	}{
		{[]byte("Line 1\r\n"), false, 8, []byte("Line 1")},
		{[]byte("Line 1\n"), false, 0, nil},
		{[]byte("Line 1"), false, 0, nil},
		{[]byte(""), false, 0, nil},
		{[]byte("Line 1\r\n"), true, 8, []byte("Line 1")},
		{[]byte("Line 1\n"), true, 7, []byte("Line 1")},
		{[]byte("Line 1"), true, 6, []byte("Line 1")},
		{[]byte(""), true, 0, nil},
	}
	for _, c := range cases {
		advance, token, _ := scanLines(c.input, c.atEOF)
		assert.Equal(t, c.advance, advance)
		assert.Equal(t, c.token, token)
	}
}

func TestStopped(t *testing.T) {
	notStoppedCh := make(chan struct{})
	isStopped := stopped(notStoppedCh)
	assert.Equal(t, false, isStopped)

	stoppedCh := make(chan struct{}, 2)
	stoppedCh <- struct{}{}
	isStopped = stopped(stoppedCh)
	assert.Equal(t, true, isStopped)
}

func TestSleepOrDone(t *testing.T) {
	doneCh := make(chan struct{}, 2)
	wait := time.Nanosecond * 20
	completed := make(chan bool)

	go func() {
		sleepOrDone(wait, doneCh)
		completed <- true
	}()

	assertReceive(t, completed, defaultTestTimeout, "sleepOrDone did not return in %v duration", wait)

	// Wait longer than timeout test because we will be returning sooner because there is a done channel message
	wait = time.Second * 5
	doneCh <- struct{}{}
	go func() {
		sleepOrDone(wait, doneCh)
		completed <- true
	}()

	assertReceive(t, completed, defaultTestTimeout, "sleepOrDone expected to return immediately but timed out")
}
