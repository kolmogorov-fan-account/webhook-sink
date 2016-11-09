package main

import (
	"net/http"
	"os"
)

func neverSayNo(w http.ResponseWriter, r *http.Request) {}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", neverSayNo)
	http.ListenAndServe(":"+port, nil)
}
