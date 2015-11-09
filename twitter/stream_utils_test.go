package twitter

import (
	"testing"

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
