package main

import (
	"fmt"
)

func main() {
	var board [8][8]rune
	specialRow := [8]rune{'r', 'k', 'b', 'q', 'k', 'b', 'k', 'r'}
	var pawnRow [8]rune
	for i := range pawnRow {
		pawnRow[i] = 'p'
	}

	for i := range board {
		for j := range board[i] {
			board[i][j] = 'o'
		}
	}

	for i := range board[0] {
		board[0][i] = specialRow[i]
		board[1][i] = pawnRow[i]
		// Substract 32 from the lowercase rune to get to the uppercase rune in ASCII
		board[7][i] = specialRow[i] - 32
		board[6][i] = pawnRow[i] - 32
	}

	for i := range board {
		for j := range board[i] {
			fmt.Printf("%-3c", board[i][j])
		}
		fmt.Printf("\n")
	}
}
