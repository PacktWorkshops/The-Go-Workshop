package main

import (
	"log"
	"net/http"
)

type hello struct{}

func(h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Hello World</h1>"
	w.Write([]byte(msg))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", hello{}))
}


