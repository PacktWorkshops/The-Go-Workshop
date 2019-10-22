package main

import "fmt"

func defineArray() [10]int {
	var arr [10]int
	return arr
}

func main() {
	fmt.Printf("%#v\n", defineArray())
}

