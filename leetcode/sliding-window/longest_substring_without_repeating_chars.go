package main

func lengthOfLongestSubstring(s string) int {
	// dynamic sliding window again?
	seen := make(map[rune]int)
	windowStart := 0
	maxLength := 0

	for windowEnd, c := range s {
		if lastSeenIndex, ok := seen[c]; ok && lastSeenIndex >= windowStart {
			windowStart = lastSeenIndex + 1
		}

		seen[c] = windowEnd

		currentLength := windowEnd - windowStart + 1
		if maxLength < currentLength {
			maxLength = currentLength
		}
	}
	return maxLength
}

func main() {
	println("3. Longest Substring Without Repeating Characters")

	// "abcabcbb"
	// 3
	println(lengthOfLongestSubstring("abcabcbb")) // 3

	// "bbbbb"
	// 1
	println(lengthOfLongestSubstring("bbbbb")) // 1

	// "pwwkew"
	// 3
	println(lengthOfLongestSubstring("pwwkew")) // 2

	// "zxyzxyz"
	// 3
	println(lengthOfLongestSubstring("zxyzxyz")) // 3

	// " "
	// 1
	println(lengthOfLongestSubstring(" ")) // 1

	// "aab"
	// 2
	println(lengthOfLongestSubstring("aab")) // 2

	// "dvdf"
	// 3
	println(lengthOfLongestSubstring("dvdf")) // 2
}
