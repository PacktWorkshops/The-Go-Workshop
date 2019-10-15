package main

import (
	"testing"
)

// TestGetDataAndReturnResponse requires the server to be running to succeed
func TestGetDataAndReturnResponse(t *testing.T) {
	data := getDataAndReturnResponse()
	if data.Message != "hello world" {
		t.Errorf("Expected string 'hello world' but received: '%s'", data)
	}
}
