package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_something(t *testing.T) {
	srv := httptest.NewServer(http.StripPrefix(
		"/statics/",
		http.FileServer(http.Dir("./public")),
	))

	rsp, err := http.Get(srv.URL + "/statics/body.css")
	if err != nil {
		t.Error(err)
	}

	expected, err := ioutil.ReadFile("./public/body.css")
	if err != nil {
		t.Error(err)
	}

	actual := make([]byte, rsp.ContentLength)
	rsp.Body.Read(actual)
	if  string(actual)!= string(expected) {
		t.Errorf("%s\n%s", string(expected),string(actual))
	}
}
