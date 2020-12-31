package main
import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)
func postFileAndReturnResponse(filename string) string {
	// create a buffer we can write the file to
	fileDataBuffer := bytes.Buffer{}
	multipartWriter := multipart.NewWriter(&fileDataBuffer)
	// open the local file we want to upload
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	// create an http formfile. This wraps our local file in a format that can be sent to the server
	formFile, err := multipartWriter.CreateFormFile("myFile", file.Name())
	if err != nil {
		log.Fatal(err)
	}
	// copy the file we want to upload into the form file wrapper
	_, err = io.Copy(formFile, file)
	if err != nil {
		log.Fatal(err)
	}
	// close the file writter. This lets it know we're done copying in data
	multipartWriter.Close()
	// create the POST request to send the file data to the server
	req, err := http.NewRequest("POST", "http://localhost:8080", &fileDataBuffer)
	if err != nil {
		log.Fatal(err)
	}
	// we set the header so the server knows about the files content
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	// send the POST request and recieve the response data
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// get data from the response body
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// return the response data
	return string(data)
}
func main() {
	// POST the file
	data := postFileAndReturnResponse("./test.txt")
	// print the response
	fmt.Println(data)
}
