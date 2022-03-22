package main

import (
	"errors"
	"fmt"
)

const rows = 9
const columns = 9

type Cell struct {
	digit int8
	fixed bool
}

type Grid [rows][columns]Cell

var (
	ErrOOBounds = errors.New("Out of Bounds")
	ErrDigit    = errors.New("Invalid Digit")
	ErrRow      = errors.New("Digit already in row")
	ErrColumn   = errors.New("Digit already in column")
	ErrRegion   = errors.New("Digital already in section")
	ErrFixedDig = errors.New("Digit is fixed")
)

func NewSudoku(digits [rows][columns]int8) *Grid {
	var grid Grid
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			d := digits[r][c]
			if d != 0 {
				grid[r][c].digit = d
				grid[r][c].fixed = true
			}
		}
	}
	return &grid
}

func (g *Grid) Set(row, column int, digit int8) error {
	switch {
	case !inBounds(row, column):
		return ErrOOBounds
	case !validDigit(digit):
		return ErrDigit
	case g.isFixed(row, column):
		return ErrFixedDig
	case g.inRow(row, digit):
		return ErrRow
	case g.inColumn(column, digit):
		return ErrColumn
	case g.inRegion(row, column, digit):
		return ErrRegion
	}
	g[row][column].digit = digit
	return nil
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows || column < 0 || column >= columns {
		return false
	}
	return true
}

func validDigit(digit int8) bool {
	return digit > 0 && digit <= 9
}

func (g *Grid) isFixed(row, column int) bool {
	return g[row][column].fixed
}

func (g *Grid) inRow(row int, digit int8) bool {
	for c := 0; c < columns; c++ {
		if g[row][c].digit == digit {
			return true
		}
	}
	return false
}

func (g *Grid) inColumn(column int, digit int8) bool {
	for r := 0; r < rows; r++ {
		if g[r][column].digit == digit {
			return true
		}
	}
	return false
}

func (g *Grid) inRegion(row, column int, digit int8) bool {
	startRow := row / 3 * 3
	startColumn := column / 3 * 3

	for r := startRow; r < startRow+3; r++ {
		for c := startColumn; c < startColumn+3; c++ {
			if g[r][c].digit == digit {
				return true
			}
		}
	}
	return false
}

func main() {

	s := NewSudoku([rows][columns]int8{
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
	// 	ErrOOBounds
	err := s.Set(0, 10, 1)
	if err != nil {
		fmt.Println(err)
	}
	// 	ErrDigit
	err = s.Set(3, 0, 10)
	if err != nil {
		fmt.Println(err)
	}
	// 	ErrRow
	err = s.Set(0, 2, 5)
	if err != nil {
		fmt.Println(err)
	}
	// 	ErrColumn
	err = s.Set(2, 0, 5)
	if err != nil {
		fmt.Println(err)
	}
	// 	ErrRegion
	err = s.Set(1, 2, 3)
	if err != nil {
		fmt.Println(err)
	}
	// 	ErrFixedDig
	err = s.Set(0, 0, 5)
	if err != nil {
		fmt.Println(err)
	}
	//	Valid
	err = s.Set(0, 2, 4)
	if err != nil {
		fmt.Println(err)
	}

	for _, row := range s {
		fmt.Println(row)
	}
}
