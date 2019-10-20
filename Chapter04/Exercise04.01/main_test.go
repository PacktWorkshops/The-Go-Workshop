package main

import "testing"

func TestDefineArray(t *testing.T) {
	if defineArray() != [10]int{} {
		t.Fail()
	}
}
