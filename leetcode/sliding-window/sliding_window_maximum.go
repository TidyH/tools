package main

import (
	"fmt"
)

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 || k > len(nums) {
		return []int{}
	}

	result := []int{}
	deque := []int{}

	for i := 0; i < len(nums); i++ {
		for len(deque) > 0 && nums[i] >= nums[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)

		if deque[0] <= i-k {
			deque = deque[1:]
		}

		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}

func main() {
	println("239. Sliding window maximum")

	// Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
	// Output: [3,3,5,5,6,7]
	ans := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3) // [3 5 5 6 7]
	fmt.Printf("%v\n", ans)
}
