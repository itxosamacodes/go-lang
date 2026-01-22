package main

import "fmt"

// func changeNum(num int) {
// 	num = 500
// 	fmt.Println("change num to : ", num)
// } // it will not change to orignal value

func changeValue(num *int) {
	*num = 500
	fmt.Println("change the value : ", *num)
}

func main() {
	num := 10

	fmt.Println("before change the value : ", num)

	changeValue(&num)

	fmt.Println("after change the value : ", num)

}
