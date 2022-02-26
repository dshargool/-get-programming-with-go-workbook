package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	width  = 80
	height = 15
)

// Define our universe with slices so we can pass it around
type Universe [][]bool

func NewUniverse() Universe {
	univ := make(Universe, height)
	for row := range univ {
		univ[row] = make([]bool, width)
	}
	return univ
}

func (u Universe) Show() {
	for row := range u {
		for _, val := range u[row] {
			if val {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func (u Universe) Seed() {
	for row := range u {
		for index := range u[row] {
			if rand.Intn(4) == 3 {
				u[row][index] = true
			}
		}
	}
}

// Alive checks the status of a cell in the universe and lets us know if it's alive or dead
// Parameters: X and Y coordinates of the cell we are looking for
// Returns: The state of the cell at the coordinates
func (u Universe) Alive(x, y int) bool {
	// Do the wrap around.  If it's within
	x = x % width
	if x < 0 {
		x = x + width
	}
	y = y % height
	if y < 0 {
		y = y + height
	}
	return u[y][x]
}

func (u Universe) Neighbours(x, y int) int {
	num_neighbours := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 || j != 0 {
				if u.Alive(x+i, y+j) {
					num_neighbours++
				}
			}
		}
	}
	return num_neighbours
}

func (u Universe) Next(x, y int) bool {
	num := u.Neighbours(x, y)
	is_alive := u.Alive(x, y)

	return num == 3 || num == 2 && is_alive
}
func Step(a, b Universe) {
	for rows := range a {
		for columns := range a[rows] {
			b[rows][columns] = a.Next(columns, rows)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	univ := NewUniverse()
	par_univ := NewUniverse()
	univ.Seed()
	for {
		univ.Show()
		//par_univ.Show()
		time.Sleep(1 * time.Second)
		Step(univ, par_univ)
		univ = par_univ
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
	}

}
