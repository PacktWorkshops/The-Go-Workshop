package main

import (
	"log"
	"fmt"
	"github.com/gorilla/mux"
)

// ExampleHandler handles the http requests send to this webserver
func 
ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
fmt.Fprintf(w, "Hello Packt")
	                          return

	log.Println("completed")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ExampleHandler)
	log.Fatal(http.ListenAndServe(":8888", r))
}
