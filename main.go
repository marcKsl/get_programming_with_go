package main

import "fmt"

func main() {
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("Exercise no 1: caesar.go")
	fmt.Println("------------------------------------------------------------------")
	m := "L fdph, L vdz, L frqtxhuhg."
	for i := 0; i < len(m); i++ {
		var b byte
		// Check if generally in range of the alphabet, lower and uppercase.
		if (m[i] >= 'a' && m[i] <= 'z') || (m[i] >= 'A' && m[i] <= 'Z') {
			if m[i] < 'A' {
				// Handle uppercase
				b = m[i] - 3
				if m[i] < 'a' {
					b += 26
				}
			} else {
				// Handle lowercase
				b = m[i] - 3
				if m[i] < 'A' {
					b += 26
				}
			}
			// Print deciphered values
			fmt.Printf("%c", b)
		} else {
			// Print non alphabetic values as they are
			fmt.Printf("%c", m[i])
		}
	}
}
