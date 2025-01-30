package main

import (
	"fmt"
)

func minEatingSpeed(piles []int, h int) int {
	low := 1
	high := 0

	for _, p := range piles {
		if p > high {
			high = p
		}
	}

	for low < high {
		totalHours := 0
		middle := low + (high-low)/2

		for _, p := range piles {
			totalHours += (p + middle - 1) / middle
		}

		if totalHours <= h {
			high = middle
		} else {
			low = middle + 1
		}
	}

	return low
}

func main() {
	fmt.Printf("875. Koko Eating Bananas\n")

	tc1 := minEatingSpeed([]int{3, 6, 7, 11}, 8)
	fmt.Printf("Starting Test #1\n")
	fmt.Printf("Expected value is: 4\n")
	fmt.Printf("Actual value is: %d\n", tc1)
	fmt.Printf(".\n")

	tc2 := minEatingSpeed([]int{30, 11, 23, 4, 20}, 5)
	fmt.Printf("Starting Test #2\n")
	fmt.Printf("Expected value is: 30\n")
	fmt.Printf("Actual value is: %d\n", tc2)
	fmt.Printf(".\n")
}
