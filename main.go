package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	gopherAmount = 29
	endurance    = 6
)

func sleepyGopher(t time.Duration, c chan int, id int) {
	time.Sleep(time.Second * t)
	fmt.Printf("Gopher %v: ...snore...\n", id)
	c <- id
}

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- v
	}
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream)
	// This will read the empty value for the type of channel
	// In this case it will read ""
	v, err := (<-downstream)
	if err == true {
		return
	}
	fmt.Println(v)
	fmt.Println(<-downstream)
}

func printGopher(upstream chan string) {
	for item := range upstream {
		fmt.Println(item)
	}
}

func main() {
	c := make(chan int)
	d := make(chan string)
	counter := 0
	go func() {
		time.Sleep(time.Second * 2)
		d <- "hello world"
	}()
	go func() {
		s := <-d
		fmt.Println(s)
	}()
	for i := 0; i < gopherAmount; i++ {
		go sleepyGopher(time.Duration(rand.Intn(endurance)), c, i)
	}
	timeout := time.After(time.Second * (endurance - 1))
loop:
	for i := 0; i < gopherAmount; i++ {
		select {
		case <-timeout:
			fmt.Println("It's over!!!")
			fmt.Printf("%v/%v gophers made it.\n", counter, gopherAmount)
			// Escape the for loop here
			break loop
		case gopherID := <-c:
			fmt.Println("gopher", gopherID, " has finished sleeping")
			counter++
			fmt.Printf("%v/%v gophers are awake now.\n", counter, gopherAmount)
		}
	}
	time.Sleep(time.Millisecond * 500)
	fmt.Println("Now the pipeline.")
	time.Sleep(time.Millisecond * 300)
	fmt.Println("------------------------------------")
	time.Sleep(time.Millisecond * 700)
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)
	time.Sleep(time.Second * 7)
	fmt.Println("The end.")
}
