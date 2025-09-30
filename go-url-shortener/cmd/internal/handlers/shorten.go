package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

var urlStore = make(map[string]string) // in-memory store

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// generate random short code
	rand.Seed(time.Now().UnixNano())
	code := randomString(6)
	urlStore[code] = req.URL

	resp := ShortenResponse{
		ShortURL: "http://localhost:8080/" + code,
	}
	json.NewEncoder(w).Encode(resp)
}

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
