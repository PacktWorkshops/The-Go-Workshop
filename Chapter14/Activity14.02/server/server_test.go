package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_posted_and_get(t *testing.T) {
	msg := Name{
		Name: "john",
	}
	bts, err := json.Marshal(msg)
	if err != nil {
		t.Error(err)
	}
	r := httptest.NewRequest(http.MethodPost, "/addName", bytes.NewReader(bts))
	r.Header.Add("Content-Type", "text/json")

	w := httptest.NewRecorder()

	AddName(w, r)

	expected, err := json.Marshal(Resp{OK: true})
	if err != nil {
		t.Error(err)
	}
	if w.Body.String() != string(expected) {
		t.Errorf("Expectes %s but recieved %s", expected, w.Body.String())
	}

	r = httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(bts))
	r.Header.Add("Content-Type", "text/json")

	w = httptest.NewRecorder()

	GetNames(w, r)
	expected, err = json.Marshal(Names{Names: []string{"john"}})
	if err != nil {
		t.Error(err)
	}
	if w.Body.String() != string(expected) {
		t.Errorf("Expectes %s but recieved %s", expected, w.Body.String())
	}
}
