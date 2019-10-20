package main

import "fmt"

func fillArray(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	return arr
}

func opArray(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	}
	return arr
}

func main() {
	var arr [10]int
	arr = fillArray(arr)
	arr = opArray(arr)
	fmt.Println(arr)
}
