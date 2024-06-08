package main

import (
	"fmt"
	"os"
	"strconv"
)

const size = 9

// isValidSudoku checks if the given board is valid
func isValidSudoku(board [size][size]int) bool {
	rows := [size][size]bool{}
	cols := [size][size]bool{}
	blocks := [size][size]bool{}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			num := board[i][j]
			if num == 0 {
				continue
			}
			num--
			blockIndex := (i/3)*3 + j/3

			if rows[i][num] || cols[j][num] || blocks[blockIndex][num] {
				return false
			}
			rows[i][num] = true
			cols[j][num] = true
			blocks[blockIndex][num] = true
		}
	}
	return true
}

// solveSudoku solves the Sudoku puzzle using backtracking
func solveSudoku(board *[size][size]int) bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == 0 {
				for num := 1; num <= size; num++ {
					board[i][j] = num
					if isValidSudoku(*board) && solveSudoku(board) {
						return true
					}
					board[i][j] = 0
				}
				return false
			}
		}
	}
	return true
}

// printBoard prints the Sudoku board
func printBoard(board [size][size]int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%d ", board[i][j])
		}
		fmt.Println()
	}
}

// parseInput parses the command line arguments into a Sudoku board
func parseInput(args []string) ([size][size]int, error) {
	var board [size][size]int

	if len(args) != size {
		return board, fmt.Errorf("invalid input")
	}

	for i, arg := range args {
		if len(arg) != size {
			return board, fmt.Errorf("invalid input")
		}
		for j := 0; j < size; j++ {
			if arg[j] == '.' {
				board[i][j] = 0
			} else {
				num, err := strconv.Atoi(string(arg[j]))
				if err != nil || num < 1 || num > size {
					return board, fmt.Errorf("invalid input")
				}
				board[i][j] = num
			}
		}
	}
	return board, nil
}

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	board, err := parseInput(os.Args[1:])
	if err != nil {
		fmt.Println("Error")
		return
	}

	if !isValidSudoku(board) {
		fmt.Println("Error")
		return
	}

	if solveSudoku(&board) {
		printBoard(board)
	} else {
		fmt.Println("Error")
	}
}
