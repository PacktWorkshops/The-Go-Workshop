package main

import (
	"testing"
)

func TestRemoveBad(t *testing.T) {
	if _, exists := getName(""); exists {
		t.Fail()
	}
	if name, exists := getName("305"); !exists || name != "Sue" {
		t.Fail()
	}
}
