package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var joinedChannel string = ""

func main() {
	// raw connection to irc.freenode
	conn, err := net.Dial("tcp", "irc.freenode.net:6667")
	if err != nil {
		log.Fatal("over")
	}
	is := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter nickname: ")
	is.Scan()
	nn := is.Text()

	// send 2 messages
	fmt.Fprintf(conn, "NICK %v\n", nn)
	fmt.Fprintf(conn, "USER guest 0 * :Coding Challenges Client\n")

	go resHandler(conn)

	for {
		is.Scan()
		err := is.Err()
		if err != nil {
			log.Fatal("scanner error")
		}

		cmd := is.Text()
		handleInpCmd(cmd, conn)
	}
}

func handleInpCmd(cmd string, conn net.Conn) {
	switch {
	case strings.HasPrefix(cmd, "/join"):
		cn := cmd[strings.Index(cmd, " ")+1:]
		fmt.Fprintf(conn, "JOIN %v\n", cn)
		joinedChannel = cn

	case strings.HasPrefix(cmd, "/part"):
		cn := cmd[strings.Index(cmd, " ")+1:]
		fmt.Fprintf(conn, "PART %v\n", cn)
		joinedChannel = ""

	case strings.HasPrefix(cmd, "/nick"):
		nn := cmd[strings.Index(cmd, " ")+1:]
		fmt.Fprintf(conn, "NICK %v\n", nn)

	case !strings.HasPrefix(cmd, "/"):
		fmt.Fprintf(conn, "PRIVMSG %v :%v\n", joinedChannel, cmd)

	case strings.HasPrefix(cmd, "/quit"):
		fmt.Fprintf(conn, "QUIT: message\n")
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
			fmt.Fprintf(conn, "PONG :%s\n", pingId)
		} else {
			origin, cmd, params := ParseMsg(resp)
			handleResCmd(origin, cmd, params)
		}

	}
}

func handleResCmd(origin string, cmd string, params []string) {
	switch cmd {
	case "NICK":
		oldNn := origin[:strings.Index(origin, "!")]
		newNn := params[0]
		fmt.Printf("\033[32m"+"MESO: %v is now known as %v\n"+"\033[0m", oldNn, newNn)

	case "PRIVMSG":
		sender := origin[:strings.Index(origin, "!")]
		fmt.Printf("\033[32m"+"%v: %v\n"+"\033[0m", sender, params[1])

	case "QUIT":
		leaver := origin[:strings.Index(origin, "!")]
		fmt.Printf("\033[32m"+"%v quit IRC. Message: %v\n"+"\033[0m", leaver, params[0])
	}
}

func ParseMsg(message string) (string, string, []string) {
	log.Println("msg: ", message)
	origin := ""
	cmd := ""
	var params []string

	tokens := strings.SplitN(message, " ", 3)
	cmd = tokens[0]
	if strings.HasPrefix(tokens[0], ":") {
		origin = tokens[0][1:]
		cmd = tokens[1]
	}

	if len(tokens) == 3 {
		if strings.Contains(tokens[2], ":") {
			trailing := tokens[2][strings.Index(tokens[2], ":")+1:]
			middle := strings.Trim(tokens[2][:strings.Index(tokens[2], ":")], " ")
			if middle != "" {
				params = strings.Split(middle, " ")
			}
			params = append(params, trailing)
		} else {
			params = strings.Split(tokens[2], " ")
		}
	}

	return origin, cmd, params
}
