package main

import "fmt"

func main() {
	var message *string

	if message == nil {
		fmt.Println("error, unexpected nil value")
		return
	}

	fmt.Println(&message)
}
