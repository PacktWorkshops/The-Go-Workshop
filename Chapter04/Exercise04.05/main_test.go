package main

import "testing"

func TestMessage(t *testing.T) {
	if message() != "It's time to Go\n" {
		t.Fail()
	}
}
