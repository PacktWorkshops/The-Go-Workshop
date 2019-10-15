package main

import (
	"encoding/json"
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
	names := Names{}
	err := json.Unmarshal(w.Body.Bytes(), &names)
	if err != nil {
		t.Error(err)
	}
	if len(names.Names) < 2 {
		t.Errorf("Expected more than 2 names. recieved: '%+v'", names)
	}
}
