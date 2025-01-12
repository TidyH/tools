package main

import "fmt"

func printGreen(str string) {
	println("\x1b[32m" + str + "\x1b[0m\n")
}

func printRed(str string) {
	println("\x1b[31m" + str + "\x1b[0m\n")
}

type testTemp struct {
	Temps    []int
	Expected []int
}

// func dailyTemperatures(temperatures []int) []int {
// 	ans := make([]int, len(temperatures))

// 	/*
// 		recursion?
// 		1. get current temp
// 		2. if next temp is greater, add 1, move to the next day
// 		3. if not, go to n+1 and add 1
// 		4. return to step 2
// 	*/

// 	for i := range temperatures {
// 		days := tempCheck(temperatures[i], i+1, temperatures)
// 		if days > 0 {
// 			ans[i] = days - i
// 		} else {
// 			ans[i] = 0
// 		}
// 	}

// 	return ans
// }

// func tempCheck(cur int, idx int, temps []int) int {
// 	if idx >= len(temps) {
// 		return 0
// 	}

// 	if temps[idx] > cur {
// 		return idx
// 	}

// 	return tempCheck(cur, idx+1, temps)
// }

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	stack := []int{}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ans[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}

	return ans
}

func main() {
	println("739. Daily Temperatures")

	tc1 := testTemp{
		Temps:    []int{73, 74, 75, 71, 69, 72, 76, 73},
		Expected: []int{1, 1, 4, 2, 1, 1, 0, 0},
	}
	tc2 := testTemp{
		Temps:    []int{30, 40, 50, 60},
		Expected: []int{1, 1, 1, 0},
	}

	testTemps := []testTemp{}

	testTemps = append(testTemps, tc1, tc2)

	for _, t := range testTemps {
		fmt.Printf("Temps to evaluate: %v\n", t.Temps)
		fmt.Printf("Expected: %v\n", t.Expected)

		results := dailyTemperatures(t.Temps)
		fmt.Printf("Actual: %v\n", results)

		if len(results) != len(t.Expected) {
			printRed("FAIL")
			break
		}
		for i, temp := range results {
			if temp != t.Expected[i] {
				printRed("FAIL")
				break
			}
		}
		printGreen("PASS")
	}

}
