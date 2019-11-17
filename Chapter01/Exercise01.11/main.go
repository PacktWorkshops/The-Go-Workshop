package main

import "fmt"

func main() {
	visits := 15
	// Equal (==)
	fmt.Println("First visit     :", visits == 1)
	// Not equal (!=)
	fmt.Println("Return visit    :", visits != 1)
	// Equal or greater than (>=) and (&&) less than (<)
	fmt.Println("Silver member   :", visits >= 10 && visits < 21)
	// Greater than (>) and (&&) equal or less than (<=)
	fmt.Println("Gold member     :", visits > 20 && visits <= 30)
	// Greater than (>)
	fmt.Println("Platinum member :", visits > 30)
}
