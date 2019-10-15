package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type PageWithCounter struct{
	Counter int `json:"views"`
	Heading string `json:"title"`
	Content string `json:"content"`
}

func(h *PageWithCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Counter++
	bts, err := json.Marshal(h)
	if err!=nil {
		w.WriteHeader(400)
		return
	}
	w.Write([]byte(bts))
}

func main() {
	hello :=PageWithCounter{Heading: "Hello World",Content:"This is the main page"}
	cha1 := PageWithCounter{Heading: "Chapter 1",Content:"This is the first chapter"}
	cha2 := PageWithCounter{Heading: "Chapter 2",Content:"This is the second chapter"}

	http.Handle("/", &hello)
	http.Handle("/chapter1", &cha1)
	http.Handle("/chapter2", &cha2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}


