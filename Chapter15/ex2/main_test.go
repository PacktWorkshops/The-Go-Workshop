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

func Test_chapter1(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	w := httptest.NewRecorder()
	chapter1(w,r)

	if w.Body.String() != "<h1>Chapter 1</h1>" {
		t.Errorf("Expected '<h1>Chapter 1</h1>' string but received: '%s'", w.Body.String())
	}
}
