package main

import "fmt"

func main() {
	age := 12
	if age >= 11 {
		fmt.Println("teenage")
	} else if age >= 18 {
		fmt.Println("adult")
	} else {
		fmt.Println("kid")
	}

	if isAdult := true; isAdult == true {
		fmt.Println("Adult")
	} else {
		fmt.Println("not Adult")
	}

	if isAdult := true; isAdult == true {
		fmt.Println("Adult")
	} else {
		fmt.Println("not Adult")
	}

	role := "admin"
	permission := true

	if role == "admin" && permission == true {
		fmt.Println("autherized")
	} else {

		fmt.Println("unautherized")

	}
}
