package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	// raw connection to irc.freenode
	conn, err := net.Dial("tcp", "irc.freenode.net:6667")
	if err != nil {
		log.Fatal("over")
	}

	// send 2 messages
	fmt.Fprintf(conn, "NICK CCClient\n")
	fmt.Fprintf(conn, "USER guest 0 * :Coding Challenges Client\n")

	// print replies
	r := bufio.NewReader(conn)
	for {
		resp, err := r.ReadString('\n')
		log.Println(resp)
		if err != nil {
			log.Fatalf("finish")
		}

		if strings.HasPrefix(resp, "PING") {
			// pring msg
			pingId := resp[strings.Index(resp, ":")+1:]
			fmt.Fprintf(conn, "PONG :%s", pingId)
		}
	}
}
