package main

import "fmt"

var defaultOffset = 10

func main() {
	offset := defaultOffset
	fmt.Println(offset)
	offset = offset + defaultOffset
	fmt.Println(offset)
}
