package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExampleHandler(t *testing.T) {
	mockRequest, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Error(err)
	}

	httpRecorder := httptest.NewRecorder()

	handlerFunction := http.HandlerFunc(ExampleHandler)

	handlerFunction.ServeHTTP(httpRecorder, mockRequest)

	if status := httpRecorder.Code; status != 200 {
		t.Errorf("invalid status code")
	}

	if httpRecorder.Body.String() != "Hello Packt" {
		t.Errorf("handler returned unexpected body")
	}
}
