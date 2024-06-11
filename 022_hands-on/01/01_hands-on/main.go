package main

import (
	"fmt"
	"io"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "root")
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func dog(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "woof woof")
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func me(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Hello Human")
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
