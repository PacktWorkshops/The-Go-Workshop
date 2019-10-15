package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthorized_ServeHTTP(t *testing.T) {
	server := server{}

	r := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	r.Header.Set("Authorization", "superSecretToken")
	w := httptest.NewRecorder()

	server.ServeHTTP(w, r)
	if w.Body.String() != "hello client!" {
		t.Errorf("Expected 'hello client!' string but received: '%s'", w.Body.String())
	}
}

func TestUnauthorized_ServeHTTP(t *testing.T) {
	server := server{}

	r := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	w := httptest.NewRecorder()

	server.ServeHTTP(w, r)
	if w.Body.String() != "Authorization token not recognized" {
		t.Errorf("Expected 'Authorization token not recognized' string but received: '%s'", w.Body.String())
	}
}
