package main

func twoSum2(nums []int, target int) []int {
	// Create a map to store numbers and their corresponding indices
	numToIndexMap := make(map[int]int)

	// Loop through the array
	for i, num := range nums {
		// Calculate the difference between the target and the current number
		diff := target - num

		// Check if the difference already exists in the map
		if idx, found := numToIndexMap[diff]; found {
			// If it exists, return the indices of the current number and the number that adds up to the target

			return []int{i, idx}
		}

		// If it doesn't exist, add the current number and its index to the map
		numToIndexMap[num] = i
	}

	// If no two numbers add up to the target, return nil
	return nil
}

func main() {
	println("1. Two Sum")

	// Local testing block
	// var nums []int = []int{2, 7, 11, 15}
	// var target int = 9

	// var nums []int = []int{3, 2, 4}
	// var target int = 6

	var nums []int = []int{2, 7, 11, 15}
	var target int = 22

	var ans []int = twoSum2(nums, target)
	println(ans[0], ans[1])
}
