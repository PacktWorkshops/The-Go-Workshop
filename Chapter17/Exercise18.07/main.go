package main

import "fmt"

// Add returns the total of two integers added together
func Add(a, b int) int {
	return a + b
}

// Multiply returns the total of one integers multipled by the other
func Multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(Add(1, 1))
	fmt.Println(Multiply(2, 2))
}
