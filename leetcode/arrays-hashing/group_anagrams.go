package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) map[string][]string {
	anagrams := make(map[string][]string)

	// turn into runes and sort
	for _, str := range strs {
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		sortedWord := string(runes)
		anagrams[sortedWord] = append(anagrams[sortedWord], str) //str appends to sorted so we have unique keys and their original as the values
	}

	return anagrams
}

func main() {
	println("49. Group Anagrams")

	var strs = []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	// Grouping anagrams
	/* Turn string into runes, sort, then create a map using the key of the sorted word,
	append the sorted word to it. This is a nice way to dynamically create keys for new values
	and add values to previous created keys.
	*/
	ans := groupAnagrams(strs)

	// character frequency
	fmt.Printf("%v\n", ans)
}
