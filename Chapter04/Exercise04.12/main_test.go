package main

import (
	"testing"
)

func TestResults(t *testing.T) {
	if l1, l2, l3 := linked(); l1 != 99 || l2 != 99 || l3 != 99 {
		t.Fail()
	}
	if nl1, nl2 := noLink(); nl1 != 99 || nl2 != 4 {
		t.Fail()
	}
	if cl1, cl2 := capLinked(); cl1 != 99 || cl2 != 99 {
		t.Fail()
	}
	if cnl1, cnl2 := capNoLink(); cnl1 != 99 || cnl2 != 4 {
		t.Fail()
	}
	if copy1, copy2, copied := copyNoLink(); copy1 != 99 || copy2 != 4 || copied != 5 {
		t.Fail()
	}
	if a1, a2 := appendNoLink(); a1 != 99 || a2 != 4 {
		t.Fail()
	}
}
