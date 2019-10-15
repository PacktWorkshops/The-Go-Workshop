package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 255; i++ {
		fmt.Printf("Decimal: %3.d Base Two: %8.b Hex:  %2.x\n", i, i, i)
	}
}

