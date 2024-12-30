package main

import "fmt"

func printGreen(str string) string {
	return "\x1b[32m" + str + "\x1b[0m\n"
}

func printRed(str string) string {
	return "\x1b[31m" + str + "\x1b[0m\n"
}

type SudokuBoard struct {
	Board    [][]string
	Expected bool
}

func (SB SudokuBoard) ConverToBoardBytes() [][]byte {
	byteBoard := make([][]byte, len(SB.Board))

	for i, row := range SB.Board {
		for _, r := range row {
			byteBoard[i] = append(byteBoard[i], r[0])
		}
	}

	return byteBoard
}

func isValidSudoku(board [][]byte) bool {
	rows := make(map[int]map[byte]int)
	cols := make(map[int]map[byte]int)
	blocks := make(map[int]map[byte]int)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]int)
		cols[i] = make(map[byte]int)
		blocks[i] = make(map[byte]int)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num == '.' {
				continue
			}

			boxIndex := 3*(i/3) + j/3

			// Check rows
			if rows[i][num]++; rows[i][num] > 1 {
				return false
			}

			// Check columns
			if cols[j][num]++; cols[j][num] > 1 {
				return false
			}

			// Check 3x3 blocks
			if blocks[boxIndex][num]++; blocks[boxIndex][num] > 1 {
				return false
			}
		}
	}

	return true
}

func main() {
	//36. Valid Sudoku
	testCases := []SudokuBoard{}

	// Valid Sudoku
	tc1 := SudokuBoard{
		Board: [][]string{
			{"5", "3", ".", ".", "7", ".", ".", ".", "."},
			{"6", ".", ".", "1", "9", "5", ".", ".", "."},
			{".", "9", "8", ".", ".", ".", ".", "6", "."},
			{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
			{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
			{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
			{".", "6", ".", ".", ".", ".", "2", "8", "."},
			{".", ".", ".", "4", "1", "9", ".", ".", "5"},
			{".", ".", ".", ".", "8", ".", ".", "7", "9"},
		},
		Expected: true,
	}

	// Invalid column
	tc2 := SudokuBoard{
		Board: [][]string{
			{"8", "3", ".", ".", "7", ".", ".", ".", "."},
			{"6", ".", ".", "1", "9", "5", ".", ".", "."},
			{".", "9", "8", ".", ".", ".", ".", "6", "."},
			{"8", ".", ".", ".", "6", ".", ".", ".", "3"}, // 8 at (0, 3) is the error
			{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
			{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
			{".", "6", ".", ".", ".", ".", "2", "8", "."},
			{".", ".", ".", "4", "1", "9", ".", ".", "5"},
			{".", ".", ".", ".", "8", ".", ".", "7", "9"},
		},
		Expected: false,
	}

	// Invalid row
	tc3 := SudokuBoard{
		Board: [][]string{
			{"8", "3", ".", ".", "7", ".", ".", ".", "."},
			{"6", ".", ".", "1", "9", "5", ".", ".", "."},
			{".", "9", "8", ".", ".", ".", ".", "6", "."},
			{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
			{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
			{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
			{".", "6", ".", ".", ".", ".", "2", "8", "."},
			{".", ".", ".", "4", "1", "9", ".", ".", "5"},
			{".", ".", ".", "9", "8", ".", ".", "7", "9"}, // 9 at (3, 8) is the error
		},
		Expected: false,
	}

	// Invalid Block
	tc4 := SudokuBoard{
		Board: [][]string{
			{"5", "3", ".", ".", "7", ".", ".", ".", "."},
			{"6", "8", ".", "1", "9", "5", ".", ".", "."}, // 8 at (1, 1) is the error
			{".", "9", "8", ".", ".", ".", ".", "6", "."},
			{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
			{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
			{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
			{".", "6", ".", ".", ".", ".", "2", "8", "."},
			{".", ".", ".", "4", "1", "9", ".", ".", "5"},
			{".", ".", ".", ".", "8", ".", ".", "7", "9"},
		},
		Expected: false,
	}

	// Combine test cases
	testCases = append(testCases, tc1, tc2, tc3, tc4)

	for i, tc := range testCases {
		fmt.Println("Running Test Case #", i+1)
		bb := tc.ConverToBoardBytes()
		result := isValidSudoku(bb)

		if result == tc.Expected {
			fmt.Println(printGreen("PASS"))
		} else {
			fmt.Println(printRed("FAIL"))
		}
	}
}
