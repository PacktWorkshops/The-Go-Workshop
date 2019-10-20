package main

import "fmt"

func message() string {
	arr := [...]string{
		"ready",
		"Get",
		"Go",
		"to",
	}
	return fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])
}

func main() {
	fmt.Print(message())
}
