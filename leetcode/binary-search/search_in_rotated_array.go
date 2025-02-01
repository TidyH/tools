package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	if len(nums) == 1 && nums[left] == target {
		return left
	}

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[left] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func main() {
	fmt.Printf("33. Search in Rotated Sorted Array\n")

	tc1 := search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	fmt.Printf("4 : %d\n", tc1)

	tc2 := search([]int{4, 5, 6, 7, 0, 1, 2}, 3)
	fmt.Printf("-1 : %d\n", tc2)

	tc3 := search([]int{1}, 0)
	fmt.Printf("-1 : %d\n", tc3)

	tc4 := search([]int{1}, 1)
	fmt.Printf("0 : %d\n", tc4)

	tc5 := search([]int{4, 5, 6, 7, 0, 1, 2}, 5)
	fmt.Printf("1 : %d\n", tc5)
}

// [4,5,6,7,0,1,2] target 5 // check if the target is higher or lower than the mid?
// mid 7; 7 is greater than 2; we set left to mid + 1; 0
