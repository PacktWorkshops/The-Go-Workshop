package main

import (
	"log"
	"net/http"
	"time"
)

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check Authorization header is what we expect
	auth := r.Header.Get("Authorization")
	if auth != "superSecretToken" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Authorization token not recognized"))
		return
	}

	// wait 10 seconds before responding
	time.Sleep(10 * time.Second)

	msg := "hello client!"
	w.Write([]byte(msg))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
