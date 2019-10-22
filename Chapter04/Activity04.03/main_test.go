package main

import (
	"reflect"
	"testing"
)

func TestRemoveBad(t *testing.T) {
	s := []string{"Good", "Good", "Good", "Good"}
	if !reflect.DeepEqual(s, removeBad()) {
		t.Fail()
	}
}
