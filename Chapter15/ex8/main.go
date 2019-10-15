package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Request struct {
	Name string
	Surname string
}

type Response struct {
	Greeting string
}

func Hello(wr http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var data Request
	err := decoder.Decode(&data)
	if err != nil {
		wr.WriteHeader(400)
		return
	}

	rsp := Response{Greeting: fmt.Sprintf("Hello %s %s", data.Name, data.Surname)}
	bts, err := json.Marshal(rsp)
	if err != nil {
		wr.WriteHeader(400)
		return
	}
	wr.Write(bts)
}

func main() {
	http.HandleFunc("/", Hello)


	log.Fatal(http.ListenAndServe(":8080", nil))
}