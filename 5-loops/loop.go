package main

import (
	"fmt"
)

func main() {
	// implemrnting while loop in the go lang using for loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	///////////////////////////////////////////

	// infinit loop
	j := 1
	for {
		fmt.Println(j)
	}

	///////////////////////////////////////////////

	// for loop
	for i := 1; i <= 10; i++ {
		fmt.Println("i love pakistan")
	}

	// continue keyword

	for i := 1; i <= 10; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}

	for i := range 4 {
		fmt.Println(i, "heloo")
	}
}
