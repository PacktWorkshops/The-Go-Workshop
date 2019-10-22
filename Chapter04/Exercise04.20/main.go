package main

import (
	"fmt"
	"math"
)

func convert() string{
	var i8 int8 = math.MaxInt8
	i := 128
	f64 := 3.14
	m := fmt.Sprintf("int8    = %v  > in64    = %v\n", i8, int64(i8))
	m += fmt.Sprintf("int     = %v  > in8     = %v\n", i, int8(i))
	m += fmt.Sprintf("int8    = %v  > float32 = %v\n", i8, float64(i8))
	m += fmt.Sprintf("float64 = %v > int     = %v\n", f64, int(f64))
	return m
}

func main() {
	fmt.Print(convert())
}
