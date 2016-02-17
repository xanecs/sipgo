package main

import (
	"fmt"
	"log"

	"github.com/xanecs/sip/sip"
)

func main() {
	conn, err := sip.Connect("10.0.0.1:5060", "fritz.box")
	if err != nil {
		log.Fatal(err)
	}

	conn.Register("622")
	fmt.Print("=======")
	sip.ReadResponse(conn)
}
