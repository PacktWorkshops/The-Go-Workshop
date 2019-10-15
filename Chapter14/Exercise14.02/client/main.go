package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type messageData struct {
	Message string `json:"message"`
}

func getDataAndReturnResponse() messageData {

	// send the GET request
	r, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	// get data from the response body
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// parse the response
	message := messageData{}
	err = json.Unmarshal(data, &message)
	if err != nil {
		log.Fatal(err)
	}

	// return the response data
	return message
}

func main() {
	data := getDataAndReturnResponse()
	fmt.Println(data.Message)
}
