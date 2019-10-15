package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostJson_ServeHTTP(t *testing.T) {
	server := server{}

	message := "{\"message\":\"hello server\"}"
	r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(message))
	w := httptest.NewRecorder()

	server.ServeHTTP(w, r)
	if w.Body.String() != message {
		t.Errorf("Expected '%s' string but received: '%s'", message, w.Body.String())
	}
}
