package main

import (
	"fmt"
)

type LCSTest struct {
	Name     string
	Numbers  []int
	Expected int
}

func printGreen(str string) string {
	return "\x1b[32m" + str + "\x1b[0m\n"
}

func printRed(str string) string {
	return "\x1b[31m" + str + "\x1b[0m\n"
}

// attempt #1 O(n log n)
// func findSequence(nums []int) int {
// 	// var sorted map[int]struct{}
// 	var sequence int = 0
// 	var tempSeq int = 0 // store temp, assign to sequence at the end

// 	fmt.Printf("%v\n", nums)
// 	sort.Ints(nums)
// 	fmt.Printf("%v\n", nums)

// 	for i := 0; i < len(nums); i++ {
// 		tempSeq++ // start at 1 for each new number -- i only need to check the failure to reset

// 		if i == len(nums)-1 {
// 			println(i, nums[i], nums[i])
// 			if nums[i] == nums[i] {
// 				if tempSeq > sequence { // if the sequence in the array is greater than our held previous sequence, overwrite
// 					println("new sequence found")
// 					sequence = tempSeq
// 				}
// 			}
// 			break
// 		}

// 		if nums[i] == nums[i+1] {
// 			tempSeq--
// 			continue
// 		}

// 		println(i, nums[i], tempSeq, sequence)
// 		if nums[i]+1 != nums[i+1] { // checking next item in the sequence so index 0 on 1,2,3,4 (1+1 = 2) equals index[1](2)
// 			if tempSeq > sequence { // if the sequence in the array is greater than our held previous sequence, overwrite
// 				println("new sequence found")
// 				sequence = tempSeq
// 			}
// 			tempSeq = 0 // reset because the sequence ended
// 		}
// 	}

// 	return sequence
// }

// O(n)
func findSequence(nums []int) int {
	seen := make(map[int]struct{})
	sequence := 0
	currentSequence := 0
	var currentNum int

	for _, n := range nums {
		seen[n] = struct{}{}
	}

	for k, _ := range seen {
		if _, ok := seen[k-1]; !ok {
			currentNum = k
			currentSequence = 1
		}

		for _, ok := seen[currentNum+1]; ok; _, ok = seen[currentNum+1] {
			currentNum++
			currentSequence++
		}

		if currentSequence > sequence {
			sequence = currentSequence
		}
	}

	return sequence
}

func main() {
	println("128. Longest Consecutive Sequence")
	/*
		Expected is the number of consecutive integers
	*/

	TestCases := []LCSTest{}

	tc1 := LCSTest{
		Name:     "Standard",
		Numbers:  []int{100, 4, 200, 1, 3, 2},
		Expected: 4,
	}

	tc2 := LCSTest{
		Name:     "Duplicate at the begining",
		Numbers:  []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
		Expected: 9,
	}

	tc3 := LCSTest{
		Name:     "Negatives",
		Numbers:  []int{-3, -2, -1, 0, 1},
		Expected: 5,
	}

	tc4 := LCSTest{
		Name:     "With Middle Duplicates",
		Numbers:  []int{1, 2, 2, 3},
		Expected: 3,
	}

	tc5 := LCSTest{
		Name:     "Empty Input",
		Numbers:  []int{},
		Expected: 0,
	}

	TestCases = append(TestCases, tc1, tc2, tc3, tc4, tc5)

	for _, tc := range TestCases {
		fmt.Println("Running Test Case #", tc.Name)
		result := findSequence(tc.Numbers)

		if result == tc.Expected {
			fmt.Println(printGreen("PASS"))

		} else {
			fmt.Println(printRed("FAIL"))
		}
		fmt.Println("Result: ", result)
		fmt.Println("Expected: ", tc.Expected)
		println("")
	}
}
