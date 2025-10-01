package main

import (
	"fmt"
	"net/http"

	"github.com/Rashkerry-newversion/golang/go-url-shortener/cmd/internal/handlers"
)

func main() {
	http.HandleFunc("/shorten", handlers.ShortenHandler)
	http.HandleFunc("/", handlers.RedirectHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
