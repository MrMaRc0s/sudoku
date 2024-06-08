package main

import (
	"fmt"
	"os"
)

const (
	size  = 9
	empty = '.'
)

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	board := make([][]byte, size)
	for i := 0; i < size; i++ {
		if len(os.Args[i+1]) != size {
			fmt.Println("Error")
			return
		}
		board[i] = []byte(os.Args[i+1])
		for _, ch := range board[i] {
			if (ch < '1' || ch > '9') && ch != empty {
				fmt.Println("Error")
				return
			}
		}
	}

	if !solveSudoku(board) {
		fmt.Println("Error")
		return
	}

	printBoard(board)
}

func solveSudoku(board [][]byte) bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == empty {
				for num := '1'; num <= '9'; num++ {
					if isValid(board, i, j, byte(num)) {
						board[i][j] = byte(num)
						if solveSudoku(board) {
							return true
						}
						board[i][j] = empty
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, num byte) bool {
	for i := 0; i < size; i++ {
		if board[row][i] == num || board[i][col] == num || board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
			return false
		}
	}
	return true
}

func printBoard(board [][]byte) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%c", board[i][j])
			if j < size-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

