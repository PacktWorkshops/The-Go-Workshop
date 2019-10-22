package main

import (
	"testing"
)

func TestGetUser(t *testing.T) {
	if _, exists := getUser(""); exists {
		t.Fail()
	}
	if name, exists := getUser("305"); !exists || name != "Sue" {
		t.Fail()
	}
}
