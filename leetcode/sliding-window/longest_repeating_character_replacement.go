package main

import "math"

func characterReplacement(s string, k int) int {
	windowStart := 0
	seen := make(map[rune]int)
	maxLength := 0

	for windowEnd, c := range s {
		seen[c]++
		delta := windowEnd - windowStart + 1

		if delta-findMaxValue(seen) <= k {
			// what do i do here? -- find the max length here?
			if delta > maxLength {
				maxLength = delta
			}
		} else { // invalid window, shift start
			seen[rune(s[windowStart])]--
			windowStart++
		}
	}

	return maxLength
}

func findMaxValue(m map[rune]int) int {
	maxValue := math.MinInt // Initialize with the smallest possible integer
	for _, value := range m {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func main() {
	println("424. Longest Repeating Character Replacement")

	// "abcdefg"
	// k = 1; expected 2
	println(characterReplacement("abcdefg", 1))

	// "abbcbb"
	// difference of most occuring and length 6 - 4 = 2;
	// diffence minus k 2 - k = 1;
	// most ocurring plus k difference 4 + 1 = 5
	// k = 1; expected 5
	println(characterReplacement("abbcbb", 1))

	// AAABABB
	// k = 1; expected 5
	println(characterReplacement("AAABABB", 1))

	// 4
	println(characterReplacement("AABABBA", 1)) // incorrectly gives 5

	// tc5 = 1
	println(characterReplacement("a", 1))
}
