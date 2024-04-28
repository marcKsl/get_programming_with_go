package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type sudokuField struct {
	grid        [9][9]int8
	initialGrid [9][9]int8
}

func (sf sudokuField) String() string {
	var s string
	for i := range sf.grid {
		for j := range sf.grid[i] {
			s += strconv.Itoa(int(sf.grid[i][j]))
			s += "  "
		}
		s += "\n"
	}
	return s
}

var (
	ErrDigitAlreadyInUse   = errors.New("Digit already in subregion, line or column.")
	ErrOutOfBounds         = errors.New("Coordinates out of bounds.")
	ErrInvalidDigit        = errors.New("Digit must be between 0-9")
	ErrForbiddenCoordinate = errors.New("Can not overwrite an initial non zero value")
)

type SudokuError []error

func (se SudokuError) Error() string {
	if se == nil {
		return ""
	}
	var s []string
	for _, err := range se {
		if err != nil && err.Error() != "" {
			s = append(s, err.Error())
		}
	}
	return strings.Join(s, "\n")
}

func newGrid(g [9][9]int8) sudokuField {
	sf := sudokuField{grid: g, initialGrid: g}
	return sf
}

func (sf sudokuField) isValidCoordinate(x, y int8) error {
	var err SudokuError
	if x > 9 || x < 0 {
		err = append(err, ErrOutOfBounds)
	}
	if y > 9 || y < 0 {
		err = append(err, ErrOutOfBounds)
	}
	if err == nil {
		if sf.initialGrid[x][y] != 0 {
			err = append(err, ErrForbiddenCoordinate)
		}
	}
	return err
}

func isValidValue(v int8) error {
	if v <= 0 || v >= 9 {
		return ErrInvalidDigit
	}
	return nil
}

func (sf *sudokuField) setDigit(x, y, v int8) error {
	var err SudokuError
	if errCoord := sf.isValidCoordinate(x, y); errCoord != nil {
		err = append(err, errCoord)
	}
	if errValue := isValidValue(v); errValue != nil {
		err = append(err, errValue)
	}
	if err != nil {
		return err
	}
	for i := range sf.grid {
		if sf.grid[i][y] == v {
			err = append(err, ErrDigitAlreadyInUse)
		}
	}
	for i := range sf.grid[x] {
		if sf.grid[x][i] == v {
			err = append(err, ErrDigitAlreadyInUse)
		}
	}
	x0 := (x / 3) * 3
	y0 := (y / 3) * 3
	if err == nil {
		for i := x0; i < x0+3; i++ {
			for j := y0; j < y0+3; j++ {
				if sf.grid[i][j] == v {
					err = append(err, ErrDigitAlreadyInUse)
				}
			}
		}
	}
	if err == nil {
		sf.grid[x][y] = v
	}
	return err
}

func (g *sudokuField) clearDigit(x, y int8) error {
	err := g.isValidCoordinate(x, y)
	if err != nil {
		return err
	}
	g.grid[x][y] = 0
	return nil
}

func main() {
	var err SudokuError
	sf := newGrid([9][9]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})
	fmt.Println(sf)
	if err := append(err, sf.setDigit(0, 2, 3)); err != nil {
		fmt.Printf("%#v", err)
	}
	// for i := range sf.grid {
	// 	for j := range sf.grid[int8(i)] {
	// 		err = append(err, sf.setDigit(int8(i), int8(j), 9))
	// 	}
	// }
	// fmt.Println(sf)
	// fmt.Println("Errors:", err)
	// fmt.Println(sf)
}
