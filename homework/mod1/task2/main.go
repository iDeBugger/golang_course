package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	go runConnToStdout(conn)
	runStdinToConn(conn)
}

func runStdinToConn(c net.Conn) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		c.Write(s.Bytes())
	}
	if err := s.Err(); err != nil {
		fmt.Println("error while scanning, err:", err)
	}
}

func runConnToStdout(c net.Conn) {
	s := bufio.NewScanner(c)
	for s.Scan() {
		os.Stdout.Write(s.Bytes())
	}
	if err := s.Err(); err != nil {
		fmt.Println("error while scanning, err:", err)
	}
}
