package main

import "fmt"

func main() {
	{
		level := "Nest 1"
		fmt.Println("Block end   :", level)
	}
	// Error: undefined: level
	fmt.Println("Main end    :", level)
}
