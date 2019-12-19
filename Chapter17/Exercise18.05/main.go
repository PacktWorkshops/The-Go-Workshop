package main

import "fmt"

func main() {
	helloString := "Hello"
	packtString := "Packt"

	jointString := fmt.Sprintf("%s", helloString, packtString)
	fmt.Println(jointString)
}
