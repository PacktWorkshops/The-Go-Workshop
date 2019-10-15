package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func getDataWithCustomOptionsAndReturnResponse() string {
	client := http.Client{Timeout: 11 * time.Second}

	// create the GET request
	req, err := http.NewRequest("POST", "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "superSecretToken")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// get data from the response body
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// return the response data
	return string(data)
}

func main() {
	data := getDataWithCustomOptionsAndReturnResponse()
	fmt.Println(data)
}
