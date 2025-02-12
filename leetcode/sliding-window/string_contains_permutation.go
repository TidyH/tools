package main

import (
	"reflect"
)

func checkInclusion(s1 string, s2 string) bool {
	windowSize := len(s1)
	seen := make(map[rune]int)
	windowStart := 0
	search := make(map[rune]int)

	for _, s := range s1 {
		search[s]++
	}

	for idx, c := range s2 {
		seen[c]++

		if idx >= windowSize-1 {
			eq := reflect.DeepEqual(search, seen)
			if eq {
				return true
			}

			seen[rune(s2[windowStart])]--
			if seen[rune(s2[windowStart])] == 0 {
				delete(seen, rune(s2[windowStart]))
			}

			windowStart++
		}
	}

	return false
}

func main() {
	println("567. Permutation in String")

	// Input: "ab"
	// String: "eidbaooo"
	// Expected: True
	println(checkInclusion("ab", "eidbaooo")) // true

	// s1:	"ab"
	// s2:	"eidboaoo"
	// ans:	false
	println(checkInclusion("ab", "eidboaoo")) // false

	// s1: "adc"
	// s2: "dcda"
	// ans: true
	println(checkInclusion("adc", "dcda"))
}
