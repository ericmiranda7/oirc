package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
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

	go resHandler(conn)

	is := bufio.NewScanner(os.Stdin)
	for {
		is.Scan()
		err := is.Err()
		if err != nil {
			log.Fatal("scanner error")
		}

		cmd := is.Text()
		handleCmd(cmd)
	}
}

func handleCmd(cmd string) {
	switch cmd {
	case "/join":
		return
	}
}

func resHandler(conn net.Conn) {
	// print replies
	// todo(): thread below
	r := bufio.NewReader(conn)
	for {
		resp, err := r.ReadString('\n')
		// log.Println(resp)
		if err != nil {
			log.Fatalf("finish")
		}

		if strings.HasPrefix(resp, "PING") {
			// pring msg
			pingId := resp[strings.Index(resp, ":")+1:]
			fmt.Fprintf(conn, "PONG :%s", pingId)
		}

		ParseMsg(resp)
	}
}

func ParseMsg(message string) (string, string, []string) {
	log.Println("msg: ", message)
	tokens := strings.Split(message, " ")
	origin := ""
	cmd := ""
	var params []string
	if strings.HasPrefix(message, ":") {
		origin = tokens[0][1:]
		cmd = tokens[1]
		params = tokens[2:]
	} else {
		cmd = tokens[0]
		params = tokens[1:]
	}

	return origin, cmd, params
}
