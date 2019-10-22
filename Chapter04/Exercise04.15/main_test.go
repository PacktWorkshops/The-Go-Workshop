package main

import (
	"reflect"
	"testing"
)

func TestGetUser(t *testing.T) {
	deleteUser("305")
	if !reflect.DeepEqual(users, map[string]string{"204": "Bob", "631": "Jake", "073": "Tracy"}) {
		t.Fail()
	}
}
