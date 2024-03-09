package main

import (
	"fmt"
	"strings"
)

func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

func main() {
	fmt.Println("Hello World!")
	worlds := []string{" Venus    ", " Earth    ", "    Mars    "}
	fmt.Println(strings.Join(worlds, "--"))
	hyperspace(worlds)
	fmt.Println(strings.Join(worlds, "--"))
}
