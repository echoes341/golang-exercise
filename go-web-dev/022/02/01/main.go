package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func check(err error) {
	if err != nil {
		log.Fatalln("Error!!!")
	}
}

func main() {
	l, err := net.Listen("tcp", ":4444")
	check(err)
	defer l.Close()

	for {
		conn, e := l.Accept()
		check(e)
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	fmt.Fprintf(conn, "Welcome. Your IP is: %s\r\n", conn.RemoteAddr())
	fmt.Fprint(conn, "Write something here: ")
	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	ln := scanner.Text()
	fmt.Fprintf(conn, "\r\nYou wrote: %s\r\n", ln)
}
