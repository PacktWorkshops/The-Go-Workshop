package main

import "testing"

func TestCompArrays(t *testing.T) {
	if comp1, comp2, _ := compArrays(); !comp1 || comp2 {
		t.Fail()
	}
}
