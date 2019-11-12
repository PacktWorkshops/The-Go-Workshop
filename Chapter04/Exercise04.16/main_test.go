package main

import (
	"testing"
)

func TestGetIDs(t *testing.T) {
	id1, id2, id3 := getIDs()
	if id1 == id2 {
		t.Fail()
	}
	if id2 != id3 {
		t.Fail()
	}
}
