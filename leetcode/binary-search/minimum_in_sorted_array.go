package main

import "fmt"

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	middle := left + (right-left)/2

	for left < right {
		mid := left + (right-left)/2

		if nums[left] < nums[right] {
			return nums[left]
		}

		if nums[middle] < nums[right] {
			// 1,2,3,4,5 -> 1,2,3
			// 4,5,1,2,3 -> 4,5,1
			right = mid // set right index to current middle
		} else if nums[middle] > nums[right] {
			// 3,4,5,1,2 -> 5,1,2
			// 2,3,4,5,1 -> 4,5,1
			left = mid + 1 // set left index to current middle
		}
	}

	return nums[left]
}

func main() {
	fmt.Printf("153. Find Minimum in Rotated Sorted Array\n")

	tc1 := findMin([]int{3, 4, 5, 1, 2})
	fmt.Printf("%d\n", tc1)

	tc2 := findMin([]int{4, 5, 6, 7, 0, 1, 2})
	fmt.Printf("%d\n", tc2)
}
