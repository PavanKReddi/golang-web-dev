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
	switch method {
	case "GET":
		handleGet(c)
	default:
		handleNotAllowed(c)
	}
}

func handleGet(c net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Code Gangsta</title>
		</head>
		<body>
			<h1>"HOLY COW THIS IS LOW LEVEL"</h1>
		</body>
		</html>
	`
	response := buildResponse("200 OK", "text/html", body)
	io.WriteString(c, response)
}

func handleNotAllowed(c net.Conn) {
	body := "Method Not Allowed\n"
	response := buildResponse("405 Method Not Allowed", "text/plain", body)
	io.WriteString(c, response)
}

func buildResponse(status, contentType, body string) string {
	return fmt.Sprintf("HTTP/1.1 %s\r\nContent-Length: %d\r\nContent-Type: %s\r\n\r\n%s",
		status, len(body), contentType, body)
}
