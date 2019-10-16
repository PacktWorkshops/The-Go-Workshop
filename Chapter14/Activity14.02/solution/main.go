package main
import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
var url = "http://localhost:8088"
type Name struct {
	Name string `json:"name"`
}
type Names struct {
	Names []string `json:"names"`
}
type Resp struct {
	OK bool `json:"ok"`
}
func addNameAndParseResponse(nameToAdd string) error {
	name := Name{Name: nameToAdd}
	nameBytes, err := json.Marshal(name)
	if err != nil {
		return err
	}
	r, err := http.Post(fmt.Sprintf("%s/addName", url), "text/json", bytes.NewReader(nameBytes))
	if err != nil {
		return err
	}
	// get data from the response body
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	resp := Resp{}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}
	if !resp.OK {
		return errors.New("response not ok")
	}
	return nil
}
func getDataAndParseResponse() []string {
	// send the GET request
	r, err := http.Get(fmt.Sprintf("%s/", url))
	if err != nil {
		log.Fatal(err)
	}
	// get data from the response body
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	names := Names{}
	err = json.Unmarshal(data, &names)
	if err != nil {
		log.Fatal(err)
	}
	// return the data
	return names.Names
}
func main() {
	err := addNameAndParseResponse("Electric")
	if err != nil {
		log.Fatal(err)
	}
	err = addNameAndParseResponse("Boogaloo")
	if err != nil {
		log.Fatal(err)
	}
	names := getDataAndParseResponse()
	for _, name := range names {
		log.Println(name)
	}
}
