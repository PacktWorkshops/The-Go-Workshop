package main

import (
	"fmt"
)

type person struct {
	lname  string
	age    int
	salary float64
}

func main() {
	fname := "Joe"
	grades := []int{100, 87, 67}
	states := map[string]string{"KY": "Kentucky", "WV": "West Virginia", "VA": "Virginia"}
	p := person{lname: "Lincoln", age: 210, salary: 25000.00}

	fmt.Printf("fname value %#v\n", fname)
	fmt.Printf("grades value %#v\n", grades)
	fmt.Printf("states value %#v\n", states)
	fmt.Printf("p value %#v\n", p)

}
