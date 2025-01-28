package main

import "fmt"

func trap(height []int) int {
	left := 0
	right := len(height) - 1
	left_max := 0
	right_max := 0
	total_water := 0

	for left < right {
		if height[left] < height[right] {
			left_max = max(left_max, height[left])
			total_water += max(0, left_max-height[left])
			left++
		} else {
			right_max = max(right_max, height[right])
			total_water += max(0, right_max-height[right])
			right--
		}
	}

	return total_water
}

func main() {
	println("Starting Next test case...")
	tc1 := trap([]int{2, 1, 0, 1, 3, 2})
	fmt.Printf("Expected: 4\n")
	fmt.Printf("Actual: %d\n", tc1)
	fmt.Printf(".\n.\n.\n")

	println("Starting Next test case...")
	tc2 := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	fmt.Printf("Expected: 6\n")
	fmt.Printf("Actual: %d\n", tc2)
	fmt.Printf(".\n.\n.\n")

	println("Starting Next test case...")
	tc3 := trap([]int{4, 2, 0, 3, 2, 5})
	fmt.Printf("Expected: 9\n")
	fmt.Printf("Actual: %d\n", tc3)
	fmt.Printf(".\n.\n.\n")

	println("Starting Next test case...")
	tc4 := trap([]int{4, 2, 3})
	fmt.Printf("Expected: 1\n")
	fmt.Printf("Actual: %d\n", tc4)
	fmt.Printf(".\n.\n.\n")

	println("Starting Next test case...")
	tc5 := trap([]int{0})
	fmt.Printf("Expected: 0\n")
	fmt.Printf("Actual: %d\n", tc5)
	fmt.Printf(".\n.\n.\n")
}
