package main

import "fmt"

type user struct {
	name    string
	age     int
	balance float64
	member  bool
}

func getUsers() []user {
	u1 := user{
		name:    "Tracy",
		age:     51,
		balance: 98.43,
		member:  true,
	}
	u2 := user{
		age:  19,
		name: "Nick",
	}
	u3 := user{
		"Bob",
		25,
		0,
		false,
	}
	var u4 user
	u4.name = "Sue"
	u4.age = 31
	u4.member = true
	u4.balance = 17.09

	return []user{u1, u2, u3, u4}
}

func main() {
	users := getUsers()
	for i := 0; i < len(users); i++ {
		fmt.Printf("%v: %#v\n", i, users[i])
	}
}
