package main

import "fmt"

// enumerated types
type orderstatus string
type Payment int

const (
	Recived  orderstatus = "Recived"
	Confirm              = "Confirm"
	Prepared             = "Prepared"
	Deliver              = "Deviverd"
)
const (
	Pay Payment = iota
	Pending
	Paid
)

func orderStatus(status orderstatus) {
	fmt.Println("The order status change to : ", status)
}
func PaymentStatus(status Payment) {
	fmt.Println("The payment status change to : ", status)
}
func main() {
	orderStatus(Recived)
	PaymentStatus(Paid)
}
