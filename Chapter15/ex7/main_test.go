package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)
func Test_posted(t *testing.T) {
	hdl, err := NewHello("./index.html")
	if err != nil {
		t.Error(err)
	}
	srv := httptest.NewServer(hdl)
	form := url.Values{}
	form.Add("name", "john")
	form.Add("surname", "smith")
	form.Add("age", "")
	req, err := http.NewRequest("POST", srv.URL, strings.NewReader(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	rsp, err := srv.Client().Do(req)
	if err != nil {
		t.Error(err)
	}
	expected, err := ioutil.ReadFile("./teststatics/filled.html")
	if err != nil {
		t.Error(err)
	}
	actual := make([]byte, rsp.ContentLength)
	rsp.Body.Read(actual)
	if  string(actual)!= string(expected) {
		t.Errorf("\n%s\n%s", string(expected),string(actual))
	}
}