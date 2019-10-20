package main

import "testing"

func TestFindLongest(t *testing.T) {
	if findLongest([]string{"Get", "ready", "to", "Go"}) != "ready" {
		t.Fail()
	}
}
