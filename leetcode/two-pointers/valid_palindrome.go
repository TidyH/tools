package main

import (
	"strings"
	"unicode"
)

type TestCase struct {
	Name   string
	Input  string
	Output bool
}

func isPalindrome(s string) bool {
	// Remove white space and non alpha-numeric characters
	var sb strings.Builder
	p := strings.ToLower(s)

	for _, r := range p {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			sb.WriteRune(r)
		}
	}

	filteredString := sb.String()
	left := 0
	right := sb.Len() - 1

	for left < right {
		if filteredString[left] != filteredString[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func main() {
	println("125. Valid Palindrome")

	cases := []TestCase{}

	tc1 := TestCase{
		Name:   "Test Case 1",
		Input:  "A man, a plan, a canal: Panama",
		Output: true,
	}

	cases = append(cases, tc1)

	for _, tc := range cases {
		println(tc.Name)
		test := isPalindrome(tc.Input)

		if test == tc.Output {
			println("PASS")
		} else {
			println("FAIL")
		}
	}

}
