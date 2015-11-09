package twitter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewExponentialBackOff(t *testing.T) {
	b := newExponentialBackOff()
	assert.Equal(t, 5*time.Second, b.InitialInterval)
	assert.Equal(t, 2.0, b.Multiplier)
	assert.Equal(t, 320*time.Second, b.MaxInterval)
}

func TestNewAggressiveExponentialBackOff(t *testing.T) {
	b := newAggressiveExponentialBackOff()
	assert.Equal(t, 1*time.Minute, b.InitialInterval)
	assert.Equal(t, 2.0, b.Multiplier)
	assert.Equal(t, 16*time.Minute, b.MaxInterval)
}
