package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func handle(c net.Conn) {
	defer c.Close()
	r := bufio.NewScanner(c)
	for r.Scan() {
		ln := r.Text()
		fmt.Println(ln)
		if ln == "" {
			break
		}
	}
	fmt.Println("Code got here.")
	io.WriteString(c, "I see you connected.")
}

func check(e error) {
	if e != nil {
		log.Fatalln("Impossible")
	}
}
func main() {
	l, e := net.Listen("tcp", ":8080")
	check(e)
	defer l.Close()
	for {
		c, e := l.Accept()
		check(e)
		go handle(c)
	}
}
