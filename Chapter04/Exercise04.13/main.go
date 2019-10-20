package main

import "fmt"

func getUsers() map[string]string {
	users := map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
	}
	users["073"] = "Tracy"
	return users
}

func main() {
	fmt.Println("Users:", getUsers())
}
