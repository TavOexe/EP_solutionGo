package main

import (
	"fmt"
)

func nqueen(n int, row int, cols, dig1, dig2 []bool, board []string, ans *[][]string) {
	if row == n {
		*ans = append(*ans, append([]string{}, board...))
		return
	}
	for col := 0; col < n; col++ {
		// guard for non placable queens
		if cols[col] || dig1[row+col] || dig2[row-col+(n-1)] {
			// if not placable
			// try another column!
			continue
		}
		cols[col] = true
		dig1[row+col] = true
		dig2[row-col+(n-1)] = true
		board[row] = board[row][:col] + "Q" + board[row][col+1:]
		// place
		// place
		// place
		// place
		// return
		nqueen(n, row+1, cols, dig1, dig2, board, ans)
		// backtracking
		// clean up
		// with this level of recursion
		cols[col] = false
		dig1[row+col] = false
		dig2[row-col+(n-1)] = false
		board[row] = board[row][:col] + board[row][col+1:]
	}

}

func solveNQueens(n int) [][]string {
	ans := [][]string{}
	// columns and diagnal directions under attack
	cols := make([]bool, n)
	// [false, false, false] where n=3
	dig1 := make([]bool, 2*n)  // Bottom Left -> Top Right: row+col
	dig2 := make([]bool, 2*n)  // Top Left -> Bottom Right: row column
	board := make([]string, n) // string representation of the board
	// initialize the board
	for i := 0; i < n; i++ {
		board[i] = string.Repeat(".", n)
	}
	nqueen(n, 0, cols, dig1, dig2, board, &ans)
	return ans

}
func main() {
	fmt.Println(solveNQueens(4))
}
