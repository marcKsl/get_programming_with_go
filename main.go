package main

import (
	"fmt"
	"time"
)

// It's a nice example, could test user input and error handling here pretty good.

type turtle struct {
	x, y int
}

type beach [][]bool

func (b *beach) showTurtle(t turtle) {
	for i := range *b {
		for j := range (*b)[i] {
			if j == t.x && i == t.y {
				fmt.Printf("*")
			} else {
				if (*b)[i][j] {
					fmt.Printf(" ")
				} else {
					fmt.Printf(".")
				}
			}
		}
		fmt.Println()
	}
	time.Sleep(time.Millisecond * 200)
	fmt.Println("\033[H")
}

func (t *turtle) move(direction int) {
	if direction > 3 {
		return
	}
	switch direction {
	case 0:
		(*t).y += 1
	case 1:
		(*t).x += 2
	case 2:
		(*t).y -= 1
	case 3:
		(*t).x -= 2
	}
}

func initialize(w, h int) beach {
	b := make([][]bool, h)
	for i := range b {
		b[i] = make([]bool, w)
	}
	return b
}

func main() {
	beach := initialize(60, 20)
	turtle := turtle{20, 10}
	for i := 0; i <= 10; i++ {
		beach.showTurtle(turtle)
		turtle.move(1)
	}
	for i := 0; i <= 3; i++ {
		beach.showTurtle(turtle)
		turtle.move(2)
	}
	for i := 0; i <= 10; i++ {
		beach.showTurtle(turtle)
		turtle.move(3)
	}
	for i := 0; i <= 3; i++ {
		beach.showTurtle(turtle)
		turtle.move(0)
	}
}
