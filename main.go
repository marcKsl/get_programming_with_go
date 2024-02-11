package main

import "fmt"

func main() {
	const year = 2400
	fmt.Println((year%4 == 0 && year%100 != 0) || year%400 == 0)

}
