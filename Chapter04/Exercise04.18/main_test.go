package main

import (
	"testing"
)

func TestGetUsers(t *testing.T) {
	if a, b := compare(); !a || !b {
		t.Fail()
	}
}
