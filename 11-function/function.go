package main

import (
	"fmt"
	"strings"
)

func add(a, b int) int {
	return a + b
}

func char(char ...string) string {
	return strings.Join(char, ",")
}
func sum(num ...int) int {
	total := 0
	for _, v := range num {
		total = v + total

	}
	return total
}

func main() {
	result := add(11, 22)
	fmt.Println(result)

	charch := char("helo", "one", "two", "three")
	fmt.Println(charch)

	nums := []int{2, 12, 35, 32, 76, 98}

	resultSum := sum(nums...)
	fmt.Println(resultSum)
}
