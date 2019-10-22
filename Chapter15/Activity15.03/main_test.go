package main
import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)
func Test_name(t *testing.T) {
	hdl, err := NewHello("./index.html")
	if err != nil {
		t.Error(err)
	}
	srv := httptest.NewServer(hdl)
	rsp, err := http.Get(srv.URL + "/?name=john")
	if err != nil {
		t.Error(err)
	}
	expected, err := ioutil.ReadFile("./teststatics/john.html")
	if err != nil {
		t.Error(err)
	}
	actual := make([]byte, rsp.ContentLength)
	rsp.Body.Read(actual)
	if  string(actual)!= string(expected) {
		t.Errorf("\n%s\n%s", string(expected),string(actual))
	}
}

func Test_anonymous(t *testing.T) {
	hdl, err := NewHello("./index.html")
	if err != nil {
		t.Error(err)
	}
	srv := httptest.NewServer(hdl)
	rsp, err := http.Get(srv.URL + "/")
	if err != nil {
		t.Error(err)
	}
	expected, err := ioutil.ReadFile("./teststatics/anonymous.html")
	if err != nil {
		t.Error(err)
	}
	actual := make([]byte, rsp.ContentLength)
	rsp.Body.Read(actual)
	if  string(actual)!= string(expected) {
		t.Errorf("\n%s\n%s", string(expected),string(actual))
	}
}