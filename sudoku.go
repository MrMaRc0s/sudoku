package main

import (
	"fmt"
	"os"
)

const N = 9

// Check if num can be placed at board[row][col]
func isValid(board [N][N]int, row, col, num int) bool {
	boxStartRow, boxStartCol := 3*(row/3), 3*(col/3)
	for i := 0; i < N; i++ {
		if board[row][i] == num || board[i][col] == num || board[boxStartRow+i/3][boxStartCol+i%3] == num {
			return false
		}
	}
	return true
}

// Solve the Sudoku puzzle using backtracking
func solveSudoku(board [N][N]int) bool {
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			if board[row][col] == 0 {
				for num := 1; num <= N; num++ {
					if isValid(board, row, col, num) {
						board[row][col] = num
						if solveSudoku(board) {
							return true
						}
						board[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

// Print the Sudoku board
func printBoard(board [N][N]int) {
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			fmt.Print(board[row][col])
			if col != N-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// Convert character to digit
func charToDigit(c byte) int {
	return int(c - '0')
}

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	var board [N][N]int
	for i := 1; i <= N; i++ {
		if len(os.Args[i]) != N {
			fmt.Println("Error")
			return
		}
		for j := 0; j < N; j++ {
			if os.Args[i][j] == '.' {
				board[i-1][j] = 0
			} else {
				num := charToDigit(os.Args[i][j])
				if num < 1 || num > 9 {
					fmt.Println("Error")
					return
				}
				board[i-1][j] = num
			}
		}
	}

	if solveSudoku(board) {
		printBoard(board)
	} else {
		fmt.Println("Error")
	}
}
