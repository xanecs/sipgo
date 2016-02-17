package sip

import (
	"bytes"
	"fmt"
	"io"
)

// Verb represents the Verb of an SIP Request (like REGISTER or INVITE)
type Verb int

// Defining SIP verbs
const (
	Register Verb = iota
	Invite
)

func (v Verb) String() string {
	switch v {
	case Register:
		return "REGISTER"
	case Invite:
		return "INVITE"
	default:
		return ""
	}
}

const protocol = "SIP/2.0"

// Request represents a SIP request from UA (Client) to Server (Registrar)
type Request struct {
	Verb    Verb
	URL     string
	Headers Headers
	Body    string
}

// Send puts the Request on to an io.Writer
func (req Request) Send(out io.Writer) {
	fmt.Fprintf(out, "%s %s %s\r\n", req.Verb, req.URL, protocol)
	for name, content := range req.Headers {
		fmt.Fprintf(out, "%s: %s\r\n", name, content)
	}
	fmt.Fprintf(out, "\r\n")
	if req.Body != "" {
		fmt.Fprintf(out, "%s\r\n", req.Body)
	}
}

func (req Request) String() string {
	var out bytes.Buffer
	req.Send(&out)
	return out.String()
}

func makeRegisterRequest(server string, domain string, user string, cSeq *int) Request {
	userString := fmt.Sprintf("sip:%s@%s", user, domain)
	req := Request{
		Verb: Register,
		URL:  fmt.Sprintf("sip:%s", domain),
		Headers: Headers{
			"From":    userString,
			"To":      userString,
			"Call-ID": user,
		}.Concat(requiredHeaders(Register, server, "", cSeq)),
	}
	fmt.Print(req)
	return req
}
