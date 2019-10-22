package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Visitor struct {
	Name string
}

type Hello struct {
	tpl *template.Template
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vl := r.URL.Query()

	cust := Visitor{}


	name, ok := vl["name"]
	if ok {
		cust.Name = strings.Join(name, ",")
	}

	h.tpl.Execute(w, cust)
}

// NewHello returns a new Hello handler
func NewHello(tplPath string) (*Hello, error){
	tmpl, err := template.ParseFiles(tplPath)
	if err != nil {
		return nil, err
	}
	return &Hello{tmpl}, nil
}

func main() {
	hello, err := NewHello("./index.html")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", hello)

	log.Fatal(http.ListenAndServe(":8080", nil))
}