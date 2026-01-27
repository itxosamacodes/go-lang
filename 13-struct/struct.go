package main

import (
	"fmt"
	"time"
)

type order = struct {
	id       string
	name     string
	status   string
	createAt time.Time
}

func statusChange(o *order) {
	o.status = "Deliverd"
}

func main() {

	myorder := order{
		id:       "1",
		name:     "banana",
		status:   "recived",
		createAt: time.Now(),
	}
	fmt.Println(myorder)
	myorder.status = "paid"
	myorder.name = "airBuds"
	statusChange(&myorder)
	fmt.Println(myorder)

	language := struct {
		name   string
		isGood bool
	}{
		"golang", true}

	fmt.Println("Languages = : ", language)

}
