package main

import (
	"fmt"
	"os"
)

func removeElement(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}

func removeAll(slice []string, element string) []string {
	for index, value := range slice {
		if value == element {
			slice = removeElement(slice, index)
			return removeAll(slice, element)
		}
	}
	return slice
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Deleting element is not passed")
		os.Exit(1)
	}
	values := []string{"Good", "Good", "Bad", "Good", "Good"}
	fmt.Println(removeAll(values, os.Args[1]))
}
