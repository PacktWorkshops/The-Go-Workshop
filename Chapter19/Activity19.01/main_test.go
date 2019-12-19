package main

import (
	"fmt"
	"os"
	"testing"
)

func TestUpdatePassword(t *testing.T) {
	db, err := getConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer tearDownDB(db)
	err = initializeDB(db)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = UpdatePassword(db, "1", "password1")
	if err != nil {
		t.Error(err)
	}
}

func TestGetPassword(t *testing.T) {
	db, err := getConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer tearDownDB(db)
	err = initializeDB(db)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	resp, err := GetPassword(db, "1")
	if err != nil {
		t.Error(err)
	}
	if string(resp) != "1234" {
		t.Error("password value is wrong")
	}
}
