package main

import "testing"

func TestGetArr(t *testing.T) {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if a != getArr() {
		t.Fail()
	}
}
