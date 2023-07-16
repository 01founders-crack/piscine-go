package piscine

import "fmt"

const BoardSize = 8

type ChessBoard [BoardSize]int

func isSafe(board *ChessBoard, row, col int) bool {
	for i := 0; i < row; i++ {
		diff := board[i] - col
		if diff == 0 || diff == row-i || diff == -(row-i) {
			return false
		}
	}
	return true
}

func solveNQueens(board *ChessBoard, row int, solutions *[][]int) {
	if row == BoardSize {
		*solutions = append(*solutions, board[:])
		return
	}

	for col := 0; col < BoardSize; col++ {
		if isSafe(board, row, col) {
			board[row] = col
			solveNQueens(board, row+1, solutions)
		}
	}
}

func FindAllSolutions() [][]int {
	var solutions [][]int
	board := &ChessBoard{}
	solveNQueens(board, 0, &solutions)
	return solutions
}

func EightQueens() {
	solutions := FindAllSolutions()
	for _, solution := range solutions {
		for _, col := range solution {
			fmt.Println(col)
		}
		fmt.Println()
	}
}
