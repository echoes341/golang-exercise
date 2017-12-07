package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func serve(c net.Conn) {
	defer c.Close()
	r := bufio.NewScanner(c)
	first := true
	var method, uri string
	for r.Scan() {
		ln := r.Text()
		if first {
			first = false
			fields := strings.Fields(ln)
			method = fields[0]
			uri = fields[1]
		}
		fmt.Println(ln)
		if ln == "" {
			break
		}
	}
	body := "<!DOCTYPE html><html><head><title>Wow</title><body>"
	switch {
	case method == "GET" && uri == "/":
		body += "<p>You asked the home page... You get the homepage...maybe you'd like to <a href=\"/apply\">apply</a></p>"
	case method == "GET" && uri == "/apply":
		body += "<p>Apply...Apply what?</p>"
		body += `<form method="post" action="/apply">
		<input type="text" name="sticazzi">
		<input type="submit" value="send"></form>`
	case method == "POST" && uri == "/apply":
		body += "<p>Oh... apply here</p>"
	}

	body += "</body></html>"
	io.WriteString(c, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
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
		go serve(c)
	}
}
