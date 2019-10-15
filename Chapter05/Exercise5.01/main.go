package main

import (
	"fmt"
)

func main() {
	printAge()
}

func printAge() {
	age := make(map[string]int)
	age["John"] = 41
	age["Celina"] = 109
	age["Micah"] = 24

	for k, v := range age {
		fmt.Printf("%s is %d and ", k, v)
		if v < 40 {
			fmt.Println("young.")
		} else if v > 40 && v <= 100 {
			fmt.Println("middle age")
		} else if v > 100 {
			fmt.Println("over a century of living.")
		}
	}
}