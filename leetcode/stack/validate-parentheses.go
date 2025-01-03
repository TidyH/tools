package main

func printGreen(str string) string {
	return "\x1b[32m" + str + "\x1b[0m\n"
}

func printRed(str string) string {
	return "\x1b[31m" + str + "\x1b[0m\n"
}

type ValidParenthesis struct {
	Name   string
	Input  string
	Output bool
}

func isValid(s string) bool {
	stack := []rune{}
	var answer bool

	for _, r := range s {
		if r == '(' {
			stack = append(stack, r)
		} else if r == ')' {
			if len(stack) == 0 {
				return false
			} else if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}

		if r == '[' {
			stack = append(stack, r)
		} else if r == ']' {
			if len(stack) == 0 {
				return false
			} else if stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}

		if r == '{' {
			stack = append(stack, r)
		} else if r == '}' {
			if len(stack) == 0 {
				return false
			} else if stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	if len(stack) == 0 {
		answer = true
	} else {
		answer = false
	}

	return answer
}

func main() {
	println("20. Valid Parentheses")

	TestCases := []ValidParenthesis{}

	tc1 := ValidParenthesis{
		Name:   "blank input",
		Input:  "",
		Output: true,
	}
	tc2 := ValidParenthesis{
		Name:   "Simple ()",
		Input:  "()",
		Output: true,
	}
	tc3 := ValidParenthesis{
		Name:   "Simple [",
		Input:  "[",
		Output: false,
	}
	tc4 := ValidParenthesis{
		Name:   "Single closing",
		Input:  "]",
		Output: false,
	}
	TestCases = append(TestCases, tc1, tc2, tc3, tc4)

	for _, t := range TestCases {
		println("Starting Test Case:", t.Name)
		testans := isValid(t.Input)

		if t.Output == testans {
			println(printGreen("PASS"))
		} else {
			println(printRed("FAIL"))
		}
	}
}

/*
	These tests only check if there is the correct amount of opening and closing parenthesis, braces, and brackets
	I need to also check properly opening and closing
*/

/*
func isValidParenthesis(s string) bool {
	var answer bool = false
	var stack int = 0

	// have to add/subtract prime from each unique rune
	for _, l := range s {
		if l == '(' {
			stack += 1
		} else if l == ')' {
			stack -= 1
		}

		if l == '[' {
			stack += 2
		} else if l == ']' {
			stack += 2
		}

		if l == '{' {
			stack += 3
		} else if l == '}' {
			stack -= 3
		}
	}
*/

// 	if stack%2 == 0 {
// 		answer = true
// 	}

// 	if s == "" {
// 		answer = false
// 	}

// 	return answer
// }

// Previous solution i had i forgot about
// func isValid(s string) bool {
// 	// counter must end at 0 and never be negative
// 	m := make(map[string]int)

// 	for _, letter := range s {
// 		switch x := letter; x {
// 		case '(':
// 			m["()"] += 1
// 		case ')':
// 			m["()"] -= 1
// 		case '[':
// 			m["[]"] += 1
// 		case ']':
// 			m["[]"] -= 1
// 		case '{':
// 			m["{}"] += 1
// 		case '}':
// 			m["{}"] -= 1
// 		default:
// 			return false
// 		}
// 	}

// 	for _, v := range m {
// 		if v != 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
