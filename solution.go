package main

import (
	"fmt"
)

func main() {
	var n int
	n = 8
	fmt.Scan(&n)
	var board [][]int
	for i := 0; i < n; i++ {
		var row []int
		for j := 0; j < n; j++ {
			var value int
			fmt.Scan(&value)
			row = append(row, value)
		}
		board = append(board, row)
	}
	var solutions int
	solutions = 0
	var solution [][]int
	solution = make([][]int, n)
	for i := 0; i < n; i++ {
		solution[i] = make([]int, n)
	}
	var solutionFound bool
	solutionFound = false
	var row int
	row = 0
	var column int
	column = 0
	var queenFound bool
	queenFound = false
	var queenFoundInRow bool
	queenFoundInRow = false
	var queenFoundInColumn bool
	queenFoundInColumn = false
	var queenFoundInDiagonal bool
	queenFoundInDiagonal = false

	for !solutionFound {
		for !queenFound {
			for !queenFoundInRow {
				for !queenFoundInColumn {
					for !queenFoundInDiagonal {
						if board[row][column] == 0 {
							solution[row][column] = 0
							queenFound = true
						} else {
							solution[row][column] = 1
							column++
							if column == n {
								column = 0
								row++
								if row == n {
									queenFoundInRow = true
									row = 0
									column = 0
								}
							}
						}
					}
				}
			}
		}
		if row == n {
			solutionFound = true
			solutions++
			row = 0
			column = 0
			queenFound = false
			queenFoundInRow = false
			queenFoundInColumn = false
			queenFoundInDiagonal = false
		}
	}
	fmt.Println(solutions)

	for i := 0; i < len(solution); i++ {
		for j := 0; j < len(solution); j++ {
			fmt.Printf("%d ", solution[i][j])
		}
		fmt.Println()
	}
}
