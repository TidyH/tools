package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return ans
}

func main() {
	println("15. 3Sum")

	tc1 := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Printf("%v\n", tc1)

	tc2 := threeSum([]int{0, 1, 1})
	fmt.Printf("%v\n", tc2)

	tc3 := threeSum([]int{0, 0, 0})
	fmt.Printf("%v\n", tc3)

	tc4 := threeSum([]int{0, 0, 0, 0})
	fmt.Printf("%v\n", tc4)
}
