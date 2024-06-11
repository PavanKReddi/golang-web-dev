package main

import (
	"io"
	"net/http"
)

type pet int

func (m pet) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
		<strong>PETS</strong><br>
		<a href="/cat">cat</a><br>
		<a href="/dog">dog</a><br>
		</body></html>`
		io.WriteString(w, body)
	case "/dog":
		io.WriteString(w, "doggy doggy doggy")
	case "/cat":
		io.WriteString(w, "kitty kitty kitty")
	}
}

func main() {
	var d pet
	http.ListenAndServe(":8080", d)
}
