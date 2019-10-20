package main

import "testing"

func TestMessage(t *testing.T) {
	if message() != `0: 1
1: 4
2: 9
3: 16
` {
		t.Fail()
	}
}
