package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHello_ServeHTTP(t *testing.T) {
	server := server{}

	r := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	w := httptest.NewRecorder()

	server.ServeHTTP(w, r)
	if w.Body.String() != "{\"message\": \"hello world\"}" {
		t.Errorf("Expected '{\"message\": \"hello world\"}' string but received: '%s'", w.Body.String())
	}
}
