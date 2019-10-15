package main

import (
	"testing"
)

// TestPostDataAndReturnResponse requires the server to be running to succeed
func TestPostDataAndReturnResponse(t *testing.T) {
	msg := messageData{Message: "Testing 123"}
	data := postDataAndReturnResponse(msg)
	if data.Message != "Testing 123" {
		t.Errorf("Expected string 'Testing 123' but received: '%s'", data)
	}
}
