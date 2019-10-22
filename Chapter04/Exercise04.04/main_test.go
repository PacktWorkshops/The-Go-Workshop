package main

import "testing"

func TestMessage(t *testing.T) {
	if message() != "Get ready to Go\n" {
		t.Fail()
	}
}
