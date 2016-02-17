package sip

import "net"

// Connection manages the TCP connection to the SIP server
type Connection struct {
	net.Conn
	Domain string
	cSeq   int
	server string
}

// Register with the server
func (c Connection) Register(user string) {
	req := makeRegisterRequest(c.server, c.Domain, user, &c.cSeq)
	req.Send(c)
}

// Connect creates a new TCP connection to a sip server
func Connect(server string, domain string) (Connection, error) {
	tcpconn, err := net.Dial("tcp", server)
	if err != nil {
		return Connection{}, err
	}
	return Connection{tcpconn, domain, 0, server}, nil
}
