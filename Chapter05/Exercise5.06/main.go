package main

import (
	"fmt"
)

func main() {

	j := 9
	x := func(i int) int {
		return i * i
	}
	fmt.Printf("The square root of %d is %d\n", j, x(j))
}
