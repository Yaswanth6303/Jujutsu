package main

import "fmt"

func findLargestNumber(nums []int) int {
	if len(nums) == 0 {
		fmt.Println("Array is Empty")
		return 0
	}

	largest := 0

	for _, value := range nums {
		if value > largest {
			largest = value
		}
	}

	return largest
}

func main() {
	nums := []int{5, 12, 3, 27, 8}
	fmt.Println("Largest Number in the array is: ", findLargestNumber(nums))
}
