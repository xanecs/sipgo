package sip

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

//Headers represents a HashMap of Headers
type Headers map[string]string

func requiredHeaders(verb Verb, server string, branch string, cSeq *int) Headers {
	if branch == "" {
		b := make([]byte, 10)
		rand.Read(b)
		branch = fmt.Sprintf("z9hG4bK_%s", hex.EncodeToString(b))
	}
	*cSeq++

	return Headers{
		"CSeq":         fmt.Sprintf("%d %s", *cSeq, verb),
		"Max-Forwards": "70",
		"Via":          fmt.Sprintf("SIP/2.0/TCP %s; branch=%s", server, branch),
	}
}

// Concat combines two sets of Headers
func (dst Headers) Concat(src Headers) Headers {
	for name, content := range src {
		dst[name] = content
	}
	return dst
}
