package main

import "fmt"

func removeIdentical(upstream, downstream chan string) {
	buf := ""
	for u := range upstream {
		if buf != u {
			buf = u
			downstream <- u
		}
	}
	close(downstream)
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	s := []string{"Hello World", "Hello Spain", "Hello Germany", "Hello GB", "Hello France", "Hello Denmark", "Hello Denmark", "Hello France"}
	go removeIdentical(c0, c1)
	go func() {
		for i := range c1 {
			fmt.Println(i)
		}
	}()
	for _, i := range s {
		c0 <- i
	}
}
