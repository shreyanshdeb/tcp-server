package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var htmlBody string

func init() {
	htmlBody = `<!DOCTYPE html><html lang="en"><head> <meta charset="UTF-8"> <meta name="viewport" content="width=device-width, initial-scale=1.0"> <meta http-equiv="X-UA-Compatible" content="ie=edge"> <title>TCP Server</title></head><body> <h1>%s</h1> <h3>%s</h3> <ul style="list-style: none;"> <li><a href="/home">Home</a></li><li><a href="/about">About</a></li><li> </ul> <form action="%s" method="POST"> <input type="submit" value="POST"> </form></li></body></html>`
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	request(conn)

	conn.Close()
}

func request(conn net.Conn) {
	var method, uri, protocol string
	scanner := bufio.NewScanner(conn)
	i := 0
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			flds := strings.Fields(ln)
			method = flds[0]
			uri = flds[1]
			protocol = flds[2]
			fmt.Printf("\nMethod: %s\nURL: %s\nProtocol: %s\n\n", method, uri, protocol)
		}
		if ln == "" {
			break
		}
		i++
	}
	response(conn, method, uri)
}

func response(conn net.Conn, method, uri string) {

	switch method {
	case "GET":
		getHandler(conn, uri)
		break
	case "POST":
		postHandler(conn, uri)
		break
	default:
		break
	}

}

func getHandler(conn net.Conn, uri string) {
	fmt.Fprintln(conn, "HTTP/1.1 200 OK\r")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(htmlBody))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, htmlBody, uri, "GET", uri)
}

func postHandler(conn net.Conn, uri string) {
	fmt.Fprintln(conn, "HTTP/1.1 200 OK\r")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(htmlBody))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, htmlBody, uri, "POST", uri)
}