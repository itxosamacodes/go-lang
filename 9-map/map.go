package main

import "fmt"

func main() {

	student := make(map[string]int)
	student["rollno"] = 12
	student["marks"] = 321

	_, ok := student["rollno"]
	if ok {
		fmt.Println("roll no exist")
	} else {
		fmt.Println("roll no dosnt exist")
	}

	delete(student, "makrs")

	fmt.Println("student after delet marks : ", student)

	clear(student)

	fmt.Println("student after clear the map  : ", student)

	// intinlize and assign value

	prduct := map[string]int{
		"quantity": 23,
		"price":    53,
		"sell":     234,
	}

	fmt.Println(prduct)

}
