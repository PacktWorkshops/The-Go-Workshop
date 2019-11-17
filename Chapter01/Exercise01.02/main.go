package main

import (
	"fmt"
)

// Package level variable
var foo string = "bar"

func main() {
	// Function level variable
	var baz string = "qux"

	fmt.Println(foo, baz)
}
