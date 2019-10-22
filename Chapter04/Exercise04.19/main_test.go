package main

import (
	"testing"
)

func TestGetUsers(t *testing.T) {
	dots :=  getDots()
	if len(dots) != 4 {
		t.Fatal()
	}
	if dots[0].name != "" || dots[1].name != "A" || dots[2].name != "B" || dots[3].name != "C"   {
		t.Fail()
	}
}
