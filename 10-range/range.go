package main

import "fmt"

func main() {
	aray := []int{1, 2, 3, 4, 4, 56, 7, 8, 2, 4, 2, 1, 345, 34}

	for i, v := range aray {
		fmt.Println("index = : ", i, "Value = : ", v)
	}

	for i := range "usama" {
		fmt.Println(i)
	}

	human := map[string]string{
		"name = ":     "usama",
		"fisrtName =": "muhammad",
		"gender =":    "male",
		"city =":      "islamabad",
	}

	for k, v := range human {
		fmt.Println(k, v)
	}
}
