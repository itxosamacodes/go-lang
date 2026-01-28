package main

import "fmt"

func displayArr(arr []int) {
	for _, item := range arr {
		fmt.Println(item)
	}
}
func main() {

	displayArr([]int{1, 2, 3, 4, 5, 5, 63, 2, 34})

}
