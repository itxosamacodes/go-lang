package main

import "fmt"

func main() {
	num := []int{}
	num = append(num, 2)
	num = append(num, 3)
	num = append(num, 4)
	fmt.Println(num)
	fmt.Println(num[0:2])

	num[2] = 80
	fmt.Println(num)
}
