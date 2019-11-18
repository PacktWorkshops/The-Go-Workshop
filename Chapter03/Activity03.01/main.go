package main

import "fmt"

func main() {
	taxTotal := .0
	// Cake
	taxTotal += salesTax(.99, .075)
	// Milk
	taxTotal += salesTax(2.75, .015)
	// Butter
	taxTotal += salesTax(.87, .02)
	// Total
	fmt.Println("Sales Tax Total: ", taxTotal)
}

func salesTax(cost float64, taxRate float64) float64 {
	return cost * taxRate
}
