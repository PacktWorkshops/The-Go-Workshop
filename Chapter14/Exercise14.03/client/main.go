package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type messageData struct {
	Message string `json:"message"`
}

func postDataAndReturnResponse(msg messageData) messageData {
	// turn the message into bytes we can POST to the server
	jsonBytes, _ := json.Marshal(msg)

	// send the POST request
	r, err := http.Post("http://localhost:8080", "application/json", bytes.NewBuffer(jsonBytes))
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
	// create the message we want to POST to the server
	msg := messageData{Message: "Hi Server!"}

	// POST the data
	data := postDataAndReturnResponse(msg)

	// print the response
	fmt.Println(data.Message)
}
