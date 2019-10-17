package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// ExampleHandler handles the http requests send to this webserver
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Packt")
	log.Println("completed")
	return
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ExampleHandler)
	log.Fatal(http.ListenAndServe(":8888", r))
}
