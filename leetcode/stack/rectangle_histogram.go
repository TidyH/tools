package main

func largestRectangleArea(heights []int) int {
	ans := 0
	stack := []int{}
	heights = append(heights, 0)

	for i, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= h {
			// Calculate area for popped bar
			poppedIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // Pop

			height := heights[poppedIndex]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			area := height * width
			if area > ans {
				ans = area
			}
		}
		stack = append(stack, i) // Push current index
	}

	return ans
}

func main() {
	println("84. Largest Rectangle in Histogram")

	println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))

	test := []int{3, 4, 1, 0, 2, 3}
	println(largestRectangleArea(test))
	println(largestRectangleArea([]int{2, 1, 2}))
}
