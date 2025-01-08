package main

import (
	"strconv"
)

func printGreen(str string) {
	println("\x1b[32m" + str + "\x1b[0m\n")
}

func printRed(str string) {
	println("\x1b[31m" + str + "\x1b[0m\n")
}

func evalRPN(tokens []string) int {
	var stack []int
	for _, t := range tokens {
		switch t {
		case "+":
			last := stack[len(stack)-1]
			secondLast := stack[len(stack)-2]
			result := last + secondLast

			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		case "-":
			last := stack[len(stack)-1]
			secondLast := stack[len(stack)-2]
			result := secondLast - last

			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		case "*":
			last := stack[len(stack)-1]
			secondLast := stack[len(stack)-2]
			result := last * secondLast

			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		case "/":
			last := stack[len(stack)-1]
			secondLast := stack[len(stack)-2]
			result := secondLast / last

			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		default:
			result, _ := strconv.Atoi(t)
			stack = append(stack, result)
		}
	}

	return stack[0] // should always be left with one
}

type RPNToken struct {
	Token  []string
	Name   string
	Output int
}

func main() {
	println("150. Evaluate Reverse Polish Notation")

	Tests := []RPNToken{}

	tc1 := RPNToken{
		Token:  []string{"2", "1", "+", "3", "*"},
		Name:   "Test 1 - Double Operation",
		Output: 9,
	}

	tc2 := RPNToken{
		Token:  []string{"4", "13", "5", "/", "+"},
		Name:   "Test 2 - Double trailing operator",
		Output: 6,
	}

	Tests = append(Tests, tc1, tc2)

	for _, tc := range Tests {
		println(tc.Name)

		results := evalRPN(tc.Token)
		println("Result: ", results)
		println("Expected Result:", tc.Output, "\n")

		if results != tc.Output {
			printRed("Test Case failed\n")
		} else {
			printGreen("Test Case passed\n")
		}
	}
}
