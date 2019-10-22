package main

import (
	"testing"
)

func TestGetUsers(t *testing.T) {
	u := user{
		name:    "Tracy",
		age:     51,
		balance: 98.43,
		member:  true,
	}
	users := getUsers()
	if len(users) != 4 || users[0] != u {
		t.Fail()
	}
}
