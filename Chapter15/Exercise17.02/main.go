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

func chapter1(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Chapter 1</h1>"
	w.Write([]byte(msg))
}

func main() {
	http.HandleFunc("/chapter1", chapter1)
	http.Handle("/", hello{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}


