package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func logon(c net.Conn, inputText string) {
	// /login?user=xx&password=pwd
	result_str := "login fail.\n"
	var (
		user string
		password string
	)

	if strings.Contains(inputText, "login"){
		pos_start := strings.Index(inputText, "user")
		if pos_start != -1 {
			pos_end := strings.Index(inputText, "&")
			if pos_end != -1 {
				user = inputText[pos_start+5:pos_end]
				fmt.Println("user:", user)
			} else {
				fmt.Fprintf(c, result_str)
				return
			}
		} else {
			fmt.Fprintf(c, result_str)
		}

		pos_start = strings.Index(inputText, "password")
		if pos_start != -1 {
			password = inputText[pos_start+9:len(inputText)]
			fmt.Println("password:", password)
			if validate_user(user, password) {
				result_str = "login ok.\n"
			}
			fmt.Fprintf(c, result_str)
		} else {
			fmt.Fprintf(c, result_str)
		}

	} else {
		fmt.Fprintf(c, result_str)
	}
}

//!+
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go logon(c, input.Text())
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
