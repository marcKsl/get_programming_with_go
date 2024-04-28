package main

import (
	"fmt"
	"net/url"
)

const adr = "https://g oogle.com"

func main() {
	fmt.Println("Input adress:", adr)
	result, e := url.Parse(adr)
	if e != nil {
		if urlErr, ok := e.(*url.Error); ok {
			fmt.Println("URL that caused the error:", urlErr.URL)
			fmt.Println("Error Op:", urlErr.Op)
			fmt.Println("Underlying error:", urlErr.Err)
		}
		fmt.Println(e)
		fmt.Printf("An error ocurred: %#v", e)
	} else {
		fmt.Println("Result:", result)
	}
}
