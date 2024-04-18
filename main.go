package main

import (
	"fmt"
	"strings"
	"time"
)

type rover struct {
}

func (r rover) talk() string {
	return "vroom vroom"
}

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type stardater interface {
	YearDay() int
	Hour() int
}

func stardate(t stardater) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}

type sol int

func (s sol) YearDay() int {
	return int(s % 668)
}

func (s sol) Hour() int {
	return 0
}

func main() {
	r := rover{}
	shout(r)

	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", stardate(day))

	s := sol(1422)
	fmt.Printf("%.1f Happy birthday\n", stardate(s))

}
