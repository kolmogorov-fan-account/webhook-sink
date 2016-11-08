package main

import (
	"net/http"
	"os"
)

func neverSayNo(w http.ResponseWriter, r *http.Request) {}

func main() {
	address := os.Getenv("ADDRESS")
	if address == "" {
		address = "127.0.0.1"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", neverSayNo)
	http.ListenAndServe(address+":"+port, nil)
}
