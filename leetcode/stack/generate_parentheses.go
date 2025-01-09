package main

import "fmt"

func printGreen(str string) {
	println("\x1b[32m" + str + "\x1b[0m\n")
}

func printRed(str string) {
	println("\x1b[31m" + str + "\x1b[0m\n")
}

func generateParenthesis(n int) []string {
	var ans []string

	generateRecusrive(&ans, "", 0, 0, n)

	return ans
}

func generateRecusrive(result *[]string, currentString string, openCount int, closeCount int, n int) {
	if openCount == n && closeCount == n {
		*result = append(*result, currentString)
	}

	if openCount < n {
		generateRecusrive(result, currentString+"(", openCount+1, closeCount, n)
	}

	if closeCount < openCount {
		generateRecusrive(result, currentString+")", openCount, closeCount+1, n)
	}

}

type GenerateTest struct {
	Name     string
	N        int
	Expected []string
}

func main() {
	Tests := []GenerateTest{}
	var outcomeSentinel int

	println("22. Generate Parentheses")

	tc1 := GenerateTest{
		Name:     "One",
		N:        1,
		Expected: []string{"()"},
	}
	tc2 := GenerateTest{
		Name:     "Two",
		N:        2,
		Expected: []string{"(())", "()()"},
	}

	Tests = append(Tests, tc1, tc2)

	for _, test := range Tests {
		outcomeSentinel = 1
		println(test.Name)
		fmt.Printf("Expected: %v\n", test.Expected)

		result := generateParenthesis(test.N)
		fmt.Printf("Actual: %v\n", result)

		for i, e := range test.Expected {
			if e != result[i] {
				outcomeSentinel = 0
			}
		}

		if outcomeSentinel == 0 {
			printRed("FAIL")
		} else {
			printGreen("PASS")
		}
	}
}
