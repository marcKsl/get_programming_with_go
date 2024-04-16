package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func NewUniverse(height int, width int) Universe {
	u := make([][]bool, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

func (u Universe) Show() {
	for i := range u {
		fmt.Printf("|")
		for j := range u[i] {
			if u[i][j] {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("|")
		fmt.Printf("\n")
	}
}

func (u Universe) Seed() {
	for i := range u {
		for j := range u[i] {
			n := rand.Intn(4) // is either 0, 1, 2, 3
			switch n {
			case 0:
				u[i][j] = false
			case 1:
				u[i][j] = false
			case 2:
				u[i][j] = false
			default:
				u[i][j] = true
			}
		}
	}
}

func (u Universe) Alive(x, y int) bool {
	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}

func (u Universe) Neighbors(x, y int) int {
	i := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if u.Alive(dx+x, dy+y) {
				i++
			}

		}
	}
	return i
}

func (u Universe) Next(x, y int) bool {
	count := u.Neighbors(x, y)
	if u.Alive(x, y) {
		switch {
		case count < 2:
			return false
		case count > 3:
			return false
		default:
			return true
		}
	} else {
		if count == 3 {
			return true
		}
		return false
	}
}

func Step(a, b Universe) {
	for i := range a {
		for j := range a[i] {
			b[i][j] = a.Next(j, i)
		}
	}
}

func (u Universe) totalElements() int {
	return len(u) * len(u[0])
}

func (u Universe) allDeadOrAlive() bool {
	count := 0
	for i := range u {
		for j := range u[i] {
			if !u.Alive(j, i) {
				count++
			}
		}
	}
	if count == 0 || count == u.totalElements() {
		return true
	}
	return false
}

func main() {
	// Conways game of life.
	fmt.Println("\033[H") // clears terminal on mac
	u := NewUniverse(height, width)
	u.Seed()
	u.Show()
	u2 := u
	for {
		time.Sleep(time.Millisecond * 500)
		Step(u, u2)
		fmt.Println("\033[H") // clears terminal on mac
		u2.Show()
		u = u2
		if u.allDeadOrAlive() {
			break
		}
	}

}
