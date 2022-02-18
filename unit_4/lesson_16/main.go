package main

import "fmt"

func print_board(board [8][8]rune) {
	for _, row := range board {
		for _, column := range row {
			if column == 0 {
				fmt.Print(" ")
			} else {
				fmt.Printf("%c ", column)
			}
		}
		fmt.Println()
	}
}

func main() {
	var planets [8]string
	planets[0] = "Mercury"
	planets[1] = "Earth"
	planets[2] = "Mars"

	earth := planets[1]
	fmt.Println(len(planets))
	fmt.Println(earth)

	lit_planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	fmt.Println(len(lit_planets))

	for i, planet := range lit_planets {
		fmt.Println(i, planet)
	}

	// Chess board experiment

	var board [8][8]rune

	// Black pieces
	board[0][0] = 'r'
	board[0][1] = 'n'
	board[0][2] = 'b'
	board[0][3] = 'q'
	board[0][4] = 'k'
	board[0][5] = 'b'
	board[0][6] = 'n'
	board[0][7] = 'r'

	// White pieces
	board[7][0] = 'R'
	board[7][1] = 'N'
	board[7][2] = 'B'
	board[7][3] = 'Q'
	board[7][4] = 'K'
	board[7][5] = 'B'
	board[7][6] = 'N'
	board[7][7] = 'R'

	// Pawns
	for i, _ := range board[0] {
		board[1][i] = 'p'
		board[6][i] = 'p'
	}

	print_board(board)
}
