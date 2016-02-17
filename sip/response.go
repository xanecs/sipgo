package sip

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

// Response represents a SIP response sent from Server to Client
type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       string
}

// ReadResponse parses an io.Reader to a Respons
func ReadResponse(r io.Reader) {
	var statusCode int
	var statusText string
	var headers = make(Headers)
	_, err := fmt.Fscanf(r, "SIP/2.0 %3d %s\n", &statusCode, &statusText)
	if err != nil {
		log.Panicln("Error parsing server response: ", err)
	}
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}
		headerLine := strings.Split(line, ": ")
		headers[headerLine[0]] = headerLine[1]
	}
	fmt.Println(headers)
}
