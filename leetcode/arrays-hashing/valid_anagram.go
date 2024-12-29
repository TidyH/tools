package main

import "fmt"

func runeMap(s string, t string) bool {
	sMap := make(map[rune]int)
	tMap := make(map[rune]int)

	if len(s) == len(t) {
		for _, r := range s {
			sMap[r]++
		}

		for _, r := range t {
			tMap[r]++
		}

	} else {
		return false
	}

	fmt.Printf("%v\n", sMap)
	fmt.Printf("%v\n", tMap)

	// compare?
	for k, v := range sMap {
		if tMap[k] != v {
			return false
		}
	}

	return true
}

func isAnagram(s string, t string) {

}

func main() {
	println("242. Valid Anagram")

	// Anagrams are jumbled words
	// cat and tac
	// golf and olfg

	response := runeMap("tomo", "moto")

	// response := isAnagram("anagram", "nagaram")

	if response == true {
		println("True")
	} else {
		println("False")
	}
}
