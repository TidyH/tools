package main

import (
	"math"
)

func minWindow(s string, t string) string {
	windowStart := 0
	minLen := math.MaxInt32
	window := make(map[rune]int)
	search := make(map[rune]int)
	matched := 0
	start := 0

	for _, c := range t {
		search[c]++
	}

	if len(t) > len(s) { // if t is ever larger than s, it's impossible
		return ""
	}

	for windowEnd := 0; windowEnd < len(s); windowEnd++ {
		window[rune(s[windowEnd])]++

		if _, ok := search[rune(s[windowEnd])]; ok && window[rune(s[windowEnd])] == search[rune(s[windowEnd])] {
			matched++
		}

		for matched == len(search) {
			currentLength := windowEnd - windowStart + 1
			if currentLength < minLen {
				minLen = currentLength
				start = windowStart
			}

			leftChar := rune(s[windowStart])
			window[leftChar]--

			if _, ok := search[leftChar]; ok && window[leftChar] < search[leftChar] {
				matched--
			}
			windowStart++
		}

	}

	if minLen == math.MaxInt32 {
		return ""
	} else {
		return s[start : start+minLen]
	}

}

func main() {
	println("76. Minimum Window Substring")

	// dynamic window
	// Input: s = "ADOBECODEBANC", t = "ABC"
	//Output: "BANC"
	println(minWindow("ADOBECODEBANC", "ABC"))

	// s = "ab", t = "b"
	// output "b"
	println(minWindow("ab", "b")) // empty
}
