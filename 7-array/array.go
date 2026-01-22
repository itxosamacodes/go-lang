package main

import (
	"fmt"
)

func main() {
	todo := [3]string{"helo", "one", "three"}
	fmt.Println(todo)
	fmt.Println("to do lenght = : ", len(todo))
	num := [4]int{}
	num[0] = 2
	num[1] = 3
	num[2] = 5
	fmt.Println("array items =:", num)

	inte := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i <= len(inte); i++ {
		fmt.Println("array items equal to : ", inte[i])
	}
}
