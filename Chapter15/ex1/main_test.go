package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHello_ServeHTTP(t *testing.T) {
	hello := hello{}

	r := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	w := httptest.NewRecorder()

	hello.ServeHTTP(w,r)
	if w.Body.String() != "<h1>Hello World</h1>" {
		t.Errorf("Expected '<h1>Hello World</h1>' string but received: '%s'", w.Body.String())
	}
}

func TestHello_ServeHTTPWithPath(t *testing.T) {
	hello := hello{}

	srv := httptest.NewServer(hello)
	rsp, err := srv.Client().Get(srv.URL)
	srv.Close()
	if err != nil {
		t.Errorf("Expected no error but received %v", err)
	}

	result := make([]byte, rsp.ContentLength)
	rsp.Body.Read(result)

	if string(result) != "<h1>Hello World</h1>" {
		t.Errorf("Expected '<h1>Hello World</h1>' string but received: '%s'", string(result))
	}

}
