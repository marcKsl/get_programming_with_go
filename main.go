package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	arr[0] = 5
	// Always prefer this combined declaration and assignment
	sliceOfArr := arr[0:2]
	// Instead of this split up version
	var anotherSlice []int
	anotherSlice = arr[0:3]
	// This will create a slice containing all values from the original array
	theLastSlice := arr[:]
	// This runs from the beginning of the original array to the specified number, the number is not included
	yetAnotherSlice := arr[:2]
	// This runs from the given number to the last element, or the length of the original array
	definetelyTheLastSlice := arr[2:]
	fmt.Println(definetelyTheLastSlice)
	fmt.Println(yetAnotherSlice)
	fmt.Println(theLastSlice)
	fmt.Println(arr)
	fmt.Println(sliceOfArr)
	fmt.Println(anotherSlice)

	// Slices can also be declared and assigned at the same time without an original array.
	// We can use the composite literal syntax for this.
	// Go will still create an underlying array for this situation and will update it's length accordingly
	mySlice := []string{"One", "Two", "Three"}
	fmt.Println(mySlice)
	myArray := [3]string{"One", "Two", "Three"}
	fmt.Println(myArray)
	// This will print [3]string to indicate an array
	fmt.Printf("myArray is of type: %T\n", myArray)
	// This will print []string to indicate a slice (of an array)
	fmt.Printf("mySlice is of type: %T\n", mySlice)
}
