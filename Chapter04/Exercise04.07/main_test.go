package main

import "testing"

func TestFillAndOpArray(t *testing.T) {
	ex := [10]int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
	if opArray(fillArray([10]int{})) != ex {
		t.Fail()
	}
}
