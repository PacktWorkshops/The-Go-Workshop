package main
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
type server struct{}
func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uploadedFile, uploadedFileHeader, err := r.FormFile("myFile")
	if err != nil {
		log.Fatal(err)
	}
	defer uploadedFile.Close()
	fileContent, err := ioutil.ReadAll(uploadedFile)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("./%s", uploadedFileHeader.Filename), fileContent, 0666)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(fmt.Sprintf("%s Uploaded!", uploadedFileHeader.Filename)))
}
func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
