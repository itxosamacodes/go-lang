package main

import (
	"fmt"
)

// Generic function that accepts only int or string slices
func displayintstr[T int | string](arr []T) {
	for _, v := range arr {
		fmt.Println("array Item qual : ", v)
	}
}

// Generic function that accepts any type
func displayAnyArray[T any](arr []T) {
	for _, v := range arr {
		fmt.Println("array Item qual : ", v)
	}
}

// Generic function with two type parameters
// T -> int or bool
// names -> string
func displayArrayandValues[T int | bool, names string](arr []T, name names) {
	for _, v := range arr {
		fmt.Println("array Item qual : ", v, name)
	}
}

// Generic struct holding a slice of any type
type anyStructofArray[T any] struct {
	getArray []T
}

func main() {

	// int slice
	displayintstr([]int{1, 2, 3, 4, 5, 56, 7, 6})

	// bool slice
	displayAnyArray([]bool{true, false})

	// string slice
	displayAnyArray([]string{
		"nouman Sb",
		"Usama",
		"Full stack Bilal",
	})

	// struct with int slice
	structAray := anyStructofArray[int]{
		getArray: []int{1, 2, 3, 4, 5, 56, 7},
	}
	fmt.Println(structAray)

	// struct with string slice
	strucAray := anyStructofArray[string]{
		getArray: []string{"a", "b", "D", "W", "q"},
	}
	fmt.Println(strucAray)

	// passing array + extra value
	displayArrayandValues(
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		"Usama",
	)
}
