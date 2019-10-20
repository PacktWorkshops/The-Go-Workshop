package main

import (
	"fmt"
	"os"
)

func getUsers() map[string]string {
	return map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
		"073": "Tracy",
	}
}

func getUser(id string) (string, bool) {
	users := getUsers()
	user, exists := users[id]
	return user, exists
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("User ID not passed")
		os.Exit(1)
	}
	userID := os.Args[1]
	name, exists := getUser(userID)
	if !exists {
		fmt.Printf("Passed user ID (%v) not found.\nUsers:\n", userID)
		for key, value := range getUsers() {
			fmt.Println("    ID:", key, "Name:", value)
		}
		os.Exit(1)
	}
	fmt.Println("Name:", name)
}
