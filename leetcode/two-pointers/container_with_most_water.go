package main

import "fmt"

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	stack := 0
	currentVolume := 0

	for left < right {
		distance := right - left // should always be positive due to for loop logic

		if height[left] <= height[right] {
			currentVolume = distance * height[left]
			left++
		} else {
			currentVolume = distance * height[right]
			right--
		}

		if currentVolume > stack {
			stack = currentVolume
		}
	}

	return stack
}

func main() {
	println("11. Container With Most Water")

	tc1 := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Printf("Expected: 49, Actual: %d\n", tc1)

	tc2 := maxArea([]int{1, 1})
	fmt.Printf("Expected: 1, Actual: %d\n", tc2)

}
