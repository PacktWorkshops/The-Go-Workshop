package main

import (
	"log"
	"net/http"
)


func main() {

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, ".static-file/index.html")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}


