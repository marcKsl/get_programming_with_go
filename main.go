package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Conways game of life with arrays, no slices

// Game rules:
// cells have eight adjacent neighbours
// Switch:
// < 2 neighbours -> dies
// > 3 neighbours -> dies
// == 3 neighbours && dead -> lives
// starts off with a random board and a fixed size
// cells can be dead or alive
// game ends if all cells live or all cells died

type cells [][]bool

func (c *cells) getLivingNeighbors(x, y int) int {
	count := 0
	height := len(*c)
	width := len((*c)[0])

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == x && dy == y {
				continue
			}

			nx := (x + dx + height) % height
			ny := (y + dy + width) % width

			if (*c)[nx][ny] {
				count++
			}

		}
	}

	return count
}

func (c *cells) makeDestiny(x, y int) {
	count := c.getLivingNeighbors(x, y)
	switch {
	case count < 2:
		(*c)[x][y] = false
	case count > 3:
		(*c)[x][y] = false
	case count == 3 && (*c)[x][y]:
		(*c)[x][y] = true
	default:
		(*c)[x][y] = true
	}
}

func initialize(width, height int) cells {
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}

	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = rand.Intn(2) == 1
		}
	}

	return grid
}

func (c cells) Print() {
	for i := range c {
		for j := range c[i] {
			if c[i][j] {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

func (c1 *cells) Next(c2 *cells) {
	for i := range *c1 {
		for j := range (*c1)[i] {
			(*c2).makeDestiny(i, j)
		}
	}
	*c1 = *c2
}

func (c *cells) endGame() bool {
	allAlive := true
	allDead := true
	for i := range *c {
		for j := range (*c)[i] {
			if (*c)[i][j] {
				allDead = false
			} else {
				allAlive = false
			}
			if !allDead && !allAlive {
				return false
			}
		}
	}
	return true
}

func main() {

	const (
		height = 10
		width  = 40
	)

	fmt.Println("\033[H")
	c1 := initialize(width, height)
	c2 := initialize(width, height)
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("\033[H")

	for {
		c1.Next(&c2)
		c1.Print()
		time.Sleep(time.Millisecond * 1000)
		fmt.Println("\033[H")
		if c1.endGame() {
			break
		}
	}

}
