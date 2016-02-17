package sip

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestString(t *testing.T) {
	minimal := Request{
		Verb: Register,
		URL:  "sip:example.org",
	}

	assert.Equal(t, minimal.String(), "REGISTER sip:example.org SIP/2.0\r\n")
}
