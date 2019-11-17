package main

import "fmt"

func main() {
	count := 5
	// Add to itself
	count += 5
	fmt.Println(count)
	// Increment by 1
	count++
	fmt.Println(count)
	// Decrement by 1
	count--
	fmt.Println(count)
	// Subtract from itself
	count -= 5
	fmt.Println(count)
	// This one works for strings
	name := "John"
	name += " Smith"
	fmt.Println("Hello,", name)
}
