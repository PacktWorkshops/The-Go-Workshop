package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	intA := math.MaxInt64
	intA = intA + 1

	bigA := big.NewInt(math.MaxInt64)
	bigA.Add(bigA, big.NewInt(1))

	fmt.Println("MaxInt64: ", math.MaxInt64)
	fmt.Println("Int     :", intA)
	fmt.Println("Big Int : ", bigA.String())
}
