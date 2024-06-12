package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	// Listen for incoming TCP connections on port 8080
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
	defer l.Close()
	log.Println("Server started on :8080")

	// Accept connections in a loop
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v\n", err)
			continue
		}
		go handleConnection(conn) // Handle each connection concurrently
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()

	scanner := bufio.NewScanner(c)
	var method, url string
	isFirstLine := true

	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			// An empty line indicates the end of the HTTP request headers
			break
		}
		if isFirstLine {
			// First line should be in the format: METHOD URL PROTOCOL
			parts := strings.Split(ln, " ")
			if len(parts) >= 2 {
				method = parts[0]
				url = parts[1]
			}
			isFirstLine = false
		}
	}

	if scanner.Err() != nil {
		log.Printf("Error reading from connection: %v\n", scanner.Err())
		return
	}

	// Log request details
	log.Printf("Received request: %s %s\n", method, url)

	// Handle different HTTP methods
	switch {
	case method == "GET" && url == "/":
		handleGet(c)
	case method == "GET" && url == "/apply":
		handleGetApply(c)
	case method == "POST" && url == "/apply":
		handlePostApply(c)
	default:
		handleDefault(c)
	}
}

func handleGet(c net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>GET INDEX</title>
		</head>
		<body>
			<h1>"GET INDEX"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
		</body>
		</html>
	`
	response := buildResponse("200 OK", "text/html", body)
	io.WriteString(c, response)
}

func handleGetApply(c net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>GET DOG</title>
		</head>
		<body>
			<h1>"GET APPLY"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
			<form action="/apply" method="POST">
			<input type="hidden" value="In my good death">
			<input type="submit" value="submit">
			</form>
		</body>
		</html>
	`
	response := buildResponse("200 OK", "text/html", body)
	io.WriteString(c, response)
}

func handlePostApply(c net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>POST APPLY</title>
		</head>
		<body>
			<h1>"POST APPLY"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
		</body>
	</html>
	`
	response := buildResponse("200 OK", "text/html", body)
	io.WriteString(c, response)
}

func handleDefault(c net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>default</title>
		</head>
		<body>
			<h1>"default"</h1>
		</body>
		</html>
	`
	response := buildResponse("200 OK", "text/html", body)
	io.WriteString(c, response)
}

func buildResponse(status, contentType, body string) string {
	return fmt.Sprintf("HTTP/1.1 %s\r\nContent-Length: %d\r\nContent-Type: %s\r\n\r\n%s",
		status, len(body), contentType, body)
}
