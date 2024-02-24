package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	// We can also store arrays in arrays.

	var board [8][8]string

	board[0][0] = "r"
	board[0][7] = "r"

	for column := range board[1] {
		board[1][column] = "p"
	}

	fmt.Print(board)

	// Or for example in a sudoku game we need a 9x9 grid.
	var sudokuBoard [9][9]int
	fmt.Println(sudokuBoard)
}
