package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Packt")
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", exampleHandler)
	log.Fatal(http.ListenAndServe(":8888", r))
}
