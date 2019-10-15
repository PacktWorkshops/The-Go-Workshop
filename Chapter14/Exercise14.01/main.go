package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func getDataAndReturnResponse() string {
	// send the GET request
	r, err := http.Get("https://www.google.com")
	if err != nil {
		log.Fatal(err)
	}

	// get data from the response body
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// return the response data
	return string(data)
}

func main() {
	data := getDataAndReturnResponse()
	log.Println(data)
}
