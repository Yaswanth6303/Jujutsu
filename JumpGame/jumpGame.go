package main

import "fmt"

func CanJump(nums []int) bool {
	maxIndex := 0

	for i, num := range nums {
		if i > maxIndex {
			return false
		}

		if i+num > maxIndex {
			maxIndex = i + num
		}
	}

	return true
}

func main() {
	test1 := []int{2, 3, 1, 1, 4}
	test2 := []int{3, 2, 1, 0, 4}

	fmt.Println("Can jump from test1:", CanJump(test1)) // Expected: true
	fmt.Println("Can jump from test2:", CanJump(test2)) // Expected: false
}
