package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)
func Test_posted(t *testing.T) {
	msg := Request{
		Name: "john",
		Surname: "smith",
	}
	bts, err := json.Marshal(msg)
	if err != nil {
		t.Error(err)
	}
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(bts))
	r.Header.Add("Content-Type", "text/json")

	w := httptest.NewRecorder()

	Hello(w,r)

	expected, err := json.Marshal(Response{Greeting:"Hello john smith"})
	if err != nil {
		t.Error(err)
	}
	if w.Body.String() != string(expected) {
		t.Errorf("Expectes %s but recieved %s", expected, w.Body.String())
	}

}
