package main

import (
	"testing"
)

func TestArea(t *testing.T) {
	if area(circle{radius: 3}) != 28.274333882308138 {
		t.Error("error calculating area for circle")
	}
	if area(rectangle{length: 4, breadth: 4}) != 16.000000 {
		t.Error("error calculating area for rectangle")
	}
}
