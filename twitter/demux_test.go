package twitter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// runDemux will start a Demux handler and channel and pass all msgs to it
// then it will close and wait for the handler to finish
func runDemux(d Demux, msgs ...interface{}) {
	done := make(chan bool)
	m := make(chan interface{})
	go func() {
		d.HandleChan(m)
		// handling stopped
		done <- true
	}()
	for i := range msgs {
		m <- msgs[i]
	}
	close(m)
	// wait for handler to finish
	<-done
}

func TestDemux_Stop(t *testing.T) {
	d := NewSwitchDemux()
	m := make(chan interface{})
	done := make(chan bool)

	go func() {
		d.HandleChan(m)
		// handling stopped
		done <- true
	}()
	// Close message stream
	close(m)

	select {
	case <-done:
		break
	case <-time.After(time.Second * 3):
		t.Errorf("Demux did not stop handling after channel closed")
	}
}

func TestDemux_Tweet(t *testing.T) {
	d := NewSwitchDemux()
	handled := 0

	d.Tweet = func(*Tweet) {
		handled++
	}

	runDemux(d, &Tweet{})

	assert.Equal(t, 1, handled)
}

func TestDemux_TweetOnly(t *testing.T) {
	d := NewSwitchDemux()
	handled := 0

	d.Tweet = func(*Tweet) {
		handled++
	}

	runDemux(d, &Tweet{}, &DirectMessage{})

	assert.Equal(t, 1, handled)
}

func TestDemux_All(t *testing.T) {
	d := NewSwitchDemux()
	allHandled := 0
	tweetsHandled := 0

	d.All = func(interface{}) {
		allHandled++
	}

	d.Tweet = func(*Tweet) {
		tweetsHandled++
	}

	runDemux(d, &Tweet{}, &DirectMessage{})

	assert.Equal(t, 2, allHandled)
	assert.Equal(t, 1, tweetsHandled)
}

func TestDemux_DM(t *testing.T) {
	d := NewSwitchDemux()
	handled := 0

	d.DM = func(*DirectMessage) {
		handled++
	}

	runDemux(d, &DirectMessage{})

	assert.Equal(t, 1, handled)
}

func TestDemux_StatusDeletion(t *testing.T) {
	d := NewSwitchDemux()
	handled := 0

	d.StatusDeletion = func(*StatusDeletion) {
		handled++
	}

	runDemux(d, &StatusDeletion{})

	assert.Equal(t, 1, handled)
}

func TestDemux_LocationDeletion(t *testing.T) {
	d := NewSwitchDemux()
	handled := 0

	d.LocationDeletion = func(*LocationDeletion) {
		handled++
	}

	runDemux(d, &LocationDeletion{})

	assert.Equal(t, 1, handled)
}

func TestDemux_StreamLimit(t *testing.T) {
	d := NewSwitchDemux()
	handled := 0

	d.StreamLimit = func(*StreamLimit) {
		handled++
	}

	runDemux(d, &StreamLimit{})

	assert.Equal(t, 1, handled)
}

func TestDemux_StatusWithheld(t *testing.T) {
	d := NewSwitchDemux()
	handled := 0

	d.StatusWithheld = func(*StatusWithheld) {
		handled++
	}

	runDemux(d, &StatusWithheld{})

	assert.Equal(t, 1, handled)
}

func TestDemux_UserWithheld(t *testing.T) {
	d := NewSwitchDemux()
	handled := 0

	d.UserWithheld = func(*UserWithheld) {
		handled++
	}

	runDemux(d, &UserWithheld{})

	assert.Equal(t, 1, handled)
}

func TestDemux_StreamDisconnect(t *testing.T) {
	d := NewSwitchDemux()
	handled := 0

	d.StreamDisconnect = func(*StreamDisconnect) {
		handled++
	}

	runDemux(d, &StreamDisconnect{})

	assert.Equal(t, 1, handled)
}

func TestDemux_Warning(t *testing.T) {
	d := NewSwitchDemux()
	handled := 0

	d.Warning = func(*StallWarning) {
		handled++
	}

	runDemux(d, &StallWarning{})

	assert.Equal(t, 1, handled)
}

func TestDemux_FriendsList(t *testing.T) {
	d := NewSwitchDemux()
	handled := 0

	d.FriendsList = func(*FriendsList) {
		handled++
	}

	runDemux(d, &FriendsList{})

	assert.Equal(t, 1, handled)
}

func TestDemux_Event(t *testing.T) {
	d := NewSwitchDemux()
	handled := 0

	d.Event = func(*Event) {
		handled++
	}

	runDemux(d, &Event{})

	assert.Equal(t, 1, handled)
}

func TestDemux_Other(t *testing.T) {
	d := NewSwitchDemux()
	handled := 0

	d.Other = func(interface{}) {
		handled++
	}

	runDemux(d, &struct{}{})

	assert.Equal(t, 1, handled)
}
