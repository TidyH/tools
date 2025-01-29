package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle - 1
		}
	}

	return -1
}

func main() {
	fmt.Printf("704. Binary Search\n")

	fmt.Printf("Test #1\n")
	tc1 := search([]int{-1, 0, 3, 5, 9, 12}, 9)
	fmt.Printf("Expected: 9 exists and is at 4\n")
	fmt.Printf("Actual: %d\n", tc1)
}
