package main

import (
	"testing"
)

func TestGetDataAndReturnResponse(t *testing.T) {
	data := getDataAndReturnResponse()
	if data == "" {
		t.Errorf("Expected non empty string but received: '%s'", data)
	}
}
