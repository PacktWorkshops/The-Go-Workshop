package main

import (
	"reflect"
	"testing"
)

func TestGetUsers(t *testing.T) {
	users := map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
		"073": "Tracy",
	}
	if !reflect.DeepEqual(getUsers(), users) {
		t.Fail()
	}
}
