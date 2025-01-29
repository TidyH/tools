package main

import "fmt"

// thiis O(m + log n)
// func searchMatrix(matrix [][]int, target int) bool {
// 	row := []int{}

// 	for i := 0; i < len(matrix); i++ {
// 		if target >= matrix[i][0] {
// 			row = matrix[i]
// 		}
// 	}

// 	left := 0
// 	right := len(row) - 1

// 	for left <= right {
// 		middle := left + (right-left)/2
// 		if row[middle] == target {
// 			return true
// 		} else if row[middle] < target {
// 			left = middle + 1
// 		} else {
// 			right = middle - 1
// 		}
// 	}

// 	return false
// }

// O(log(m * n)) -- optimal
func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	left := 0
	right := rows*cols - 1

	for left <= right {
		middle := left + (right-left)/2
		row := middle / cols
		col := middle % cols

		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return false
}

func main() {
	fmt.Printf("74. Search a 2d Matrix\n")

	tc1 := searchMatrix([][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 60}},
		3)

	fmt.Printf("tc1\n")
	fmt.Printf("Expected answer is true\n")
	fmt.Printf("Actual answer is: %t", tc1)

}
