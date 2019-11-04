package main
import (
	"fmt"
)
func main() {
	for i := 1; i <= 15; i++ {
		n, s := fizzBuzz(i)
		fmt.Printf("Results:  %d %s\n", n, s)
	}
}
func fizzBuzz(i int) (int, string) {
	switch {
	case i%15 == 0:
		return i, "FizzBuzz"
	case i%3 == 0:
		return i, "Fizz"
	case i%5 == 0:
		return i, "Buzz"
	}
	return i, ""
}
