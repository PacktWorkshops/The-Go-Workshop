package main
import (
	"encoding/json"
	"log"
	"net/http"
)
var names []string
type Name struct {
	Name string `json:"name"`
}
type Names struct {
	Names []string `json:"names"`
}
type Resp struct {
	OK bool `json:"ok"`
}
func GetNames(wr http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		wr.WriteHeader(400)
		return
	}
	rsp := Names{Names: names}
	bts, err := json.Marshal(rsp)
	if err != nil {
		wr.WriteHeader(500)
		return
	}
	wr.Write(bts)
}
func AddName(wr http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		wr.WriteHeader(400)
		return
	}
	decoder := json.NewDecoder(req.Body)
	var data Name
	err := decoder.Decode(&data)
	if err != nil {
		wr.WriteHeader(400)
		return
	}
	names = append(names, data.Name)
	rsp := Resp{OK: true}
	bts, err := json.Marshal(rsp)
	if err != nil {
		wr.WriteHeader(400)
		return
	}
	wr.Write(bts)
}
func main() {
	http.HandleFunc("/addName", AddName)
	http.HandleFunc("/", GetNames)
	log.Fatal(http.ListenAndServe(":8088", nil))
}
