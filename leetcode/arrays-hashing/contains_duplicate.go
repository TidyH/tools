package main

func checkDuplicate(nums []int) bool {
	seen := make(map[int]struct{})

	for _, num := range nums {
		if _, exists := seen[num]; exists {
			return true
		}
		seen[num] = struct{}{}
	}

	return false
}

func main() {
	// nums := []int{1, 2, 3, 4, 5}
	nums := []int{1, 1, 3, 4, 10, 20, 10, 22}

	checkDuplicate(nums)
}
