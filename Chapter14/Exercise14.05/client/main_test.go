package main

import (
	"testing"
)

// TestGetDataWithCustomOptionsAndReturnResponse requires the server to be running to succeed
func TestGetDataWithCustomOptionsAndReturnResponse(t *testing.T) {
	data := getDataWithCustomOptionsAndReturnResponse()
	if data != "hello client!" {
		t.Errorf("Expected string 'hello client!' but received: '%s'", data)
	}
}
