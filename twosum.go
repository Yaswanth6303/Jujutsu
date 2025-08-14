package main

import "fmt"

// TwoSum finds two indices in nums array that add up to target
// Returns the indices of the two numbers that add up to target
func TwoSum(nums []int, target int) []int {
	// Create a map to store number and its index
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		// Check if complement exists in the map
		if index, found := numMap[complement]; found {
			return []int{index, i}
		}

		// Store the current number and its index
		numMap[num] = i
	}

	// No solution found
	return []int{}
}

func main() {
	// Test cases
	fmt.Println("Two Sum Algorithm")
	fmt.Println("=================")

	// Test case 1
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	result1 := TwoSum(nums1, target1)
	fmt.Printf("Input: nums = %v, target = %d\n", nums1, target1)
	fmt.Printf("Output: %v\n", result1)
	fmt.Printf("Explanation: nums[%d] + nums[%d] = %d + %d = %d\n\n",
		result1[0], result1[1], nums1[result1[0]], nums1[result1[1]], target1)

	// Test case 2
	nums2 := []int{3, 2, 4}
	target2 := 6
	result2 := TwoSum(nums2, target2)
	fmt.Printf("Input: nums = %v, target = %d\n", nums2, target2)
	fmt.Printf("Output: %v\n", result2)
	fmt.Printf("Explanation: nums[%d] + nums[%d] = %d + %d = %d\n\n",
		result2[0], result2[1], nums2[result2[0]], nums2[result2[1]], target2)

	// Test case 3
	nums3 := []int{3, 3}
	target3 := 6
	result3 := TwoSum(nums3, target3)
	fmt.Printf("Input: nums = %v, target = %d\n", nums3, target3)
	fmt.Printf("Output: %v\n", result3)
	fmt.Printf("Explanation: nums[%d] + nums[%d] = %d + %d = %d\n\n",
		result3[0], result3[1], nums3[result3[0]], nums3[result3[1]], target3)
}
