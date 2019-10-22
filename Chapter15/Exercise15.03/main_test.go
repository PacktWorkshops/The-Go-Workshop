package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHello(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, `/?name="john"`, strings.NewReader(""))
	w := httptest.NewRecorder()

	Hello(w,r)

	exp :=  `Hello "john"`

	if w.Body.String() != exp {
		t.Errorf("Expected '%s' but received '%s'", exp, w.Body.String())
	}

	if w.Code != http.StatusOK {
		t.Errorf("We expected a status code of %d but we received %d", http.StatusOK, w.Code)
	}

}

func TestHello_Fail(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	w := httptest.NewRecorder()

	Hello(w,r)

	exp :=  `Missing name`

	if w.Body.String() != exp {
		t.Errorf("Expected '%s' but received '%s'", exp, w.Body.String())
	}

	if w.Code != http.StatusBadRequest {
		t.Errorf("We expected a status code of %d but we received %d", http.StatusBadRequest, w.Code)
	}

}