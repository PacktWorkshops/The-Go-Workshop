package main

import "fmt"

func main() {
	nums := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}
	fmt.Println("Before:", nums)
	for swapped := true; swapped; {
		swapped = false
		for i := 1; i < len(nums); i++ {
			if nums[i-1] > nums[i] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
				swapped = true
			}
		}
	}
	fmt.Println("After :", nums)
}
