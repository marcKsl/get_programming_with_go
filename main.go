package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martian struct{}

func (m martian) talk() string {
	return "yip yip"
}

type laser struct{}

func (l *laser) talk() string {
	return "pew pew"
}

func main() {
	m := martian{}
	shout(m)
	shout(&m)

	l := laser{}
	// this does not work,
	// because the receiver of the talk() method on laser is of type pointer
	// laser...pointer, badumm
	// shout(l)

	// This works fine, because it satisfies the receiver type
	shout(&l)
}
