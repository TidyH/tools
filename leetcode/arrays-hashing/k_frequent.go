package main

import (
	"fmt"
	"reflect"
	"sort"
)

type NumFreq struct {
	Num  int
	Freq int
}

func findFreq(nums []int, k int) []int {
	/* Misread the quiz, i need to grab the top K items appearing most frequently
	 */
	freqMap := make(map[int]int)

	// turn into map we can evaluate
	for _, n := range nums {
		freqMap[n]++
	}

	// Collect map into struct
	NumFreqs := []NumFreq{}
	for num, freq := range freqMap {
		NumFreqs = append(NumFreqs, NumFreq{Num: num, Freq: freq})
	}

	// Sort by frequency
	sort.Slice(NumFreqs, func(i, j int) bool {
		return NumFreqs[i].Freq > NumFreqs[j].Freq
	})

	// Extract top k elements
	kFreq := []int{}
	for i := 0; i < k && i < len(NumFreqs); i++ {
		kFreq = append(kFreq, NumFreqs[i].Num)
	}

	return kFreq
}

// Test cases
type TestCase struct {
	Numbers  []int
	K        int
	Expected []int
}

func main() {
	println("347. Top K Frequent Elements")
	println("Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.")

	var TestCases = []TestCase{}
	tc1 := TestCase{
		Numbers:  []int{1, 1, 1, 2, 2, 3},
		K:        2,
		Expected: []int{1, 2},
	}
	tc2 := TestCase{
		Numbers:  []int{1, 2},
		K:        2,
		Expected: []int{1, 2},
	}
	tc3 := TestCase{
		Numbers:  []int{3, 0, 1, 0},
		K:        1,
		Expected: []int{0},
	}

	TestCases = append(TestCases, tc1, tc2, tc3)

	for _, tc := range TestCases {
		ans := findFreq(tc.Numbers, tc.K)
		fmt.Printf("Running Testcase: %v\n", tc.Numbers)
		if reflect.DeepEqual(ans, tc.Expected) { // check if equal
			println("PASS")
			fmt.Printf("Answer: %v\n", ans)
			fmt.Printf("Expected: %v\n", tc.Expected)
		} else {
			println("FAIL")
			fmt.Printf("Answer: %v\n", ans)
			fmt.Printf("Expected: %v\n", tc.Expected)
		}
	}
}
