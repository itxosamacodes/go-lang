package main

import (
	"fmt"
	"time"
)

type order = struct {
	name     string
	amount   float32
	status   string
	createAt time.Time
	customer
}
type customer = struct {
	custName string
	phone    string
	city     string
}

func main() {
	newOrder := order{
		name:     "AirBuds",
		amount:   33.2,
		status:   "paid",
		createAt: time.Now(),
		customer: customer{
			custName: "usama",
			phone:    "123456",
			city:     "Lahore",
		},
	}
	fmt.Println("order data = : ", newOrder)
	fmt.Println("customer data = : ", newOrder.customer)

}
