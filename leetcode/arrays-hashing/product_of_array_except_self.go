package main

import (
	"fmt"
	"reflect"
)

type TestCase struct {
	Numbers  []int
	Expected []int
}

// Attempt 1 O(n^2)
// func findProducts(nums []int) []int {
// 	var answers []int
// 	var outerCount int

// 	for len(answers) < len(nums) {
// 		var temp int = 1
// 		for i, n := range nums {
// 			if outerCount == i {
// 				continue
// 			} else {
// 				temp = temp * n
// 			}
// 		}
// 		outerCount++
// 		answers = append(answers, temp)
// 	}

// 	return answers
// }

// Prefix/Suffix Method O(n)
func findProducts(nums []int) []int {
	answers := make([]int, len(nums))
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))

	for i, _ := range nums {
		if i == 0 {
			prefix[i] = 1
		} else {
			prefix[i] = prefix[i-1] * nums[i-1]
		}
	}

	// go through in reverse order now for suffix
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			suffix[i] = 1
		} else {
			suffix[i] = suffix[i+1] * nums[i+1]
		}

		// Calculate answers finding newly found suffix -- optimized
		// if i == 0 {
		// 	answers[i] = suffix[i+1]
		// } else if i == len(nums)-1 {
		// 	answers[i] = prefix[i-1]
		// } else {
		// 	answers[i] = prefix[i-1] * suffix[i+1]
		// }

		answers[i] = prefix[i] * suffix[i]
	}

	// Find answers -- easy way
	// for i, _ := range nums {
	// 	ans := prefix[i] * suffix[i]
	// 	answers = append(answers, ans)
	// }

	return answers
}

func main() {
	fmt.Println("238. Product of Array Except Self")

	// Given an integer array nums,
	//return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
	var TestCases = []TestCase{}
	tc1 := TestCase{
		Numbers:  []int{1, 2, 3, 4},
		Expected: []int{24, 12, 8, 6},
	}
	tc2 := TestCase{
		Numbers:  []int{-1, 1, 0, -3, 3},
		Expected: []int{0, 0, 9, 0, 0},
	}

	TestCases = append(TestCases, tc1, tc2)

	// Testing the function
	for _, tc := range TestCases {
		fmt.Printf("Test Case: %v\n", tc.Numbers)
		var actual []int = findProducts(tc.Numbers)
		if reflect.DeepEqual(actual, tc.Expected) {
			fmt.Printf(printGreen("PASS"))
			fmt.Printf("Actual: %v\n", actual)
			fmt.Printf("Expected: %v\n", tc.Expected)
		} else {
			fmt.Printf(printRed("FAIL"))
			fmt.Printf("Actual: %v\n", actual)
			fmt.Printf("Expected: %v\n", tc.Expected)
		}
	}
}
