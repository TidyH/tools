package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		if numbers[left]+numbers[right] < target {
			left++
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			return []int{left + 1, right + 1}
		}
	}

	return nil
}

func main() {
	println("167. Two Sum II - Input Array Is Sorted")

	ans := twoSum([]int{7, 2, 11, 15}, 9)
	fmt.Printf("%v\n", ans)

	ans2 := twoSum([]int{2, 3, 4}, 6)
	fmt.Printf("%v\n", ans2)
}
