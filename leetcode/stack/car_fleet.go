package main

import (
	"fmt"
	"sort"
)

type FleetTest struct {
	Target   int
	Position []int
	Speed    []int
	Expected int
	Name     string
}

/*
Take remaing distance target-position = distance to go
then divide that by speed. If the next car will catch up or pass before the target, that forms a fleet
*/

type Car struct {
	Position int
	Speed    int
}

func largestRectangleArea(heights []int) int {
	ans := 0
	stack := []int{} // Stack to store indices

	// Add a sentinel -1 at the end to simplify calculations
	heights = append(heights, 0)

	for i, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1 >= h {
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
	println("853. Car Fleet")

	tests := []FleetTest{}

	tc1 := FleetTest{
		Target:   12,
		Position: []int{10, 8, 0, 5, 3},
		Speed:    []int{2, 4, 1, 1, 3},
		Expected: 3,
		Name:     "ThreeFleets",
	}

	tc2 := FleetTest{
		Target:   10,
		Position: []int{3},
		Speed:    []int{3},
		Expected: 1,
		Name:     "SingleEntry",
	}

	tc3 := FleetTest{
		Target:   10,
		Position: []int{0, 4, 2},
		Speed:    []int{2, 1, 3},
		Expected: 1,
		Name:     "Failed",
	}

	tests = append(tests, tc1, tc2, tc3)

	for _, t := range tests {
		println("Name:", t.Name)
		println("Expected:", t.Expected)
		actual := carFleet(t.Target, t.Position, t.Speed)
		println("Actual:", actual)

		if t.Expected == actual {
			fmt.Printf("Pass\n---\n")
		} else {
			fmt.Printf("Fail\n---\n")
		}
	}
}
