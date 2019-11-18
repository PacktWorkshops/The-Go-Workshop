package main

import "fmt"

func main() {
	username := "Sir_King_Über"

	// Length of a string
	fmt.Println("Bytes:", len(username))
	fmt.Println("Runes:", len([]rune(username)))

	// Limit to 10 characters
	fmt.Println(string(username[:10]))
	fmt.Println(string([]rune(username)[:10]))
}
