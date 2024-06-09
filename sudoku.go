package main

import (
	"fmt"
	"os"
	"strconv"
)

const size = 9

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

	if !hasUniqueSolution(board) {
		fmt.Println("Error")
		return
	}

	if solveSudoku(&board) {
		printBoard(board)
	} else {
		fmt.Println("Error")
	}
}

// parseInput parses the command line arguments into a Sudoku board
// checking if the input in the arguments is valid
func parseInput(raws []string) ([size][size]int, error) {
	var board [size][size]int
	if len(raws) != size {
		return board, fmt.Errorf("invalid input")
	}
	for args, raw := range raws {
		if len(raw) != size {
			return board, fmt.Errorf("invalid input")
		}
		for chars := 0; chars < size; chars++ {
			if raw[chars] == '.' {
				board[args][chars] = 0
			} else {
				num, err := strconv.Atoi(string(raw[chars]))
				if err != nil || num < 1 || num > size {
					return board, fmt.Errorf("invalid input")
				}
				board[args][chars] = num
			}
		}
	}
	return board, nil
}

// isValidSudoku checks if the given board is valid for a sudoku board
func isValidSudoku(board [size][size]int) bool {
	rows := [size][size]bool{}
	cols := [size][size]bool{}
	blocks := [size][size]bool{}
	for rawin := 0; rawin < size; rawin++ { // rawin variable is for raw index in the sudoku cells
		for colin := 0; colin < size; colin++ { // colin variable is for column index in the sudoku cells
			num := board[rawin][colin]
			if num == 0 {
				continue
			}
			num--
			blockIndex := (rawin/3)*3 + colin/3 // calculates which 3x3 sub-grid the current cell belongs to
			if rows[rawin][num] || cols[colin][num] || blocks[blockIndex][num] {
				return false
			}
			rows[rawin][num] = true
			cols[colin][num] = true
			blocks[blockIndex][num] = true
		}
	}
	return true
}

// hasUniqueSolution checks if the Sudoku puzzle has a unique solution
func hasUniqueSolution(board [size][size]int) bool {
	solutions := 0
	solveAndCountSolutions(&board, &solutions)
	return solutions == 1
}

// solveAndCountSolutions solves the Sudoku puzzle and counts the number of solutions
func solveAndCountSolutions(board *[size][size]int, solutions *int) bool {
	for rawin := 0; rawin < size; rawin++ {
		for colin := 0; colin < size; colin++ {
			if board[rawin][colin] == 0 {
				for num := 1; num <= size; num++ {
					if isValidMove(board, rawin, colin, num) {
						board[rawin][colin] = num
						if solveAndCountSolutions(board, solutions) {
							if *solutions > 1 {
								return false
							}
						}
						board[rawin][colin] = 0
					}
				}
				return false
			}
		}
	}
	*solutions++
	return true
}

// isValidMove checks if placing num in board[row][col] is valid
func isValidMove(board *[size][size]int, row, col, num int) bool {
	for rawin := 0; rawin < size; rawin++ {
		if board[row][rawin] == num || board[rawin][col] == num || board[row-row%3+rawin/3][col-col%3+rawin%3] == num {
			return false
		}
	}
	return true
}

// Sudoku solution
func solveSudoku(board *[size][size]int) bool {
	for rawin := 0; rawin < size; rawin++ {
		for colin := 0; colin < size; colin++ {
			if board[rawin][colin] == 0 {
				for num := 1; num <= size; num++ {
					board[rawin][colin] = num
					if isValidSudoku(*board) && solveSudoku(board) {
						return true
					}
					board[rawin][colin] = 0
				}
				return false
			}
		}
	}
	return true
}

// printBoard prints the Sudoku board
func printBoard(board [size][size]int) {
	for rawin := 0; rawin < size; rawin++ {
		for colin := 0; colin < size; colin++ {
			fmt.Printf("%d ", board[rawin][colin])
		}
		fmt.Println()
	}
}
