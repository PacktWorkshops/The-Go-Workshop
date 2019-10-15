package main

import (
	"testing"
)

// TestPostDataAndReturnResponse requires the server to be running to succeed
func TestPostFileAndReturnResponse(t *testing.T) {
	data := postFileAndReturnResponse("./test.txt")
	if data != "./test.txt Uploaded!" {
		t.Errorf("Expected string './test.txt Uploaded!' but received: '%s'", data)
	}
}
