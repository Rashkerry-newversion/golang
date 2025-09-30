package main

import (
	"fmt"
	"net/http"

	"go-url-shortener/cmd/internal/handlers"
)

func main() {
	http.HandleFunc("/shorten", handlers.ShortenHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
