package main

import "testing"

func TestMessage(t *testing.T) {
	m := `First   : 1 [1] [1]
Last    : 9 [9] [9]
First 5 : [1 2 3 4 5]
Last 4  : [6 7 8 9]
Middle 5: [3 4 5 6 7]
`
	if m != message() {
		t.Fail()
	}
}
