package main

import (
	"fmt"
	"sort"
)

type number struct {
	value int
	valid bool
}

func newNumber(i int) number {
	return number{value: i, valid: true}
}

func (n number) String() string {
	if n.valid == false {
		return "not set"
	}
	return fmt.Sprintf("%d", n.value)
}

type person struct {
	age int
}

func (p *person) Birthday() {
	if p == nil {
		return
	}
	p.age++
}

func sortStrings(s []string, less func(a, b int) bool) {
	if less == nil {
		less = func(a, b int) bool { return s[a] < s[b] }
	}
	sort.Slice(s, less)
}

func mirepoix(ingredients []string) []string {
	return append(ingredients, "onion", "celery", "carrot")
}

func main() {
	fmt.Println("nil")
	var nowhere *int
	fmt.Println(nowhere)
	fmt.Printf("%T\n", nowhere)
	// This would panic, because the pointer is a nil pointer, that can't be dereferenced
	// fmt.Println(*nowhere)
	// It can be caught tho:
	if nowhere != nil {
		fmt.Println(*nowhere)
	} else {
		fmt.Println("nowhere is nil, can't be dereferenced.")
	}
	// method will not crash because it guards against nil dereference
	var nobody *person
	fmt.Println(nobody)
	nobody.Birthday()

	// variables with the function type have the value nil when declared
	var fn func(a, b int) int
	fmt.Println(fn == nil)

	stringSlice := []string{"Hund", "Ulbricht", "Azeroth"}
	fmt.Println(stringSlice)
	// This will panic, because the Slice function needs a valid function but accepts nil
	// sort.Slice(stringSlice, nil)
	// So we provide a wrapper function to catch nil values
	// and provide a valid less function as default
	sortStrings(stringSlice, nil)
	fmt.Println(stringSlice)
	sortStrings(stringSlice, func(a, b int) bool { return len(stringSlice[a]) < len(stringSlice[b]) })
	fmt.Println(stringSlice)

	// slices that are declared without a composite literal or the make built-in have a nil value
	var soup []string
	fmt.Println(soup == nil)
	// range keyword and len and append built-in do catch nil
	fmt.Println(len(soup))
	for _, ingredient := range soup {
		fmt.Println(ingredient)
	}
	soup = append(soup, "onion", "celery", "carrot")
	fmt.Println(soup)

	var sopa []string
	fmt.Println(sopa)
	sopaNueva := mirepoix(sopa)
	fmt.Println(sopaNueva)
	// maps with standard declaration are also nil,
	// can be read, but panic if trying to write to it
	var suppe map[string]int
	fmt.Println(suppe == nil)

	measurement, ok := suppe["Karotten"]
	if ok {
		fmt.Println(measurement)
	}

	for ingredient, measurement := range suppe {
		fmt.Println(ingredient, measurement)
	}

	var v interface{}
	fmt.Printf("type: %T value: %v is_nil: %v\n", v, v, v == nil)
	var p *int
	v = p
	fmt.Printf("type: %T value: %v is_nil: %v\n", v, v, v == nil)
	// Shorthand to get both, the type and the value
	fmt.Printf("%#v\n", v)
	var s fmt.Stringer
	fmt.Println(s == nil)

	// An approach to avoid using nil pointers
	numberOne := newNumber(12)
	fmt.Println(numberOne)
	numberTwo := number{}
	fmt.Println(numberTwo)

}
