package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}
type payment struct {
	gatWay paymenter
}

func (p payment) pay(amount float32) {
	p.gatWay.pay(amount)
}

type jazCash struct {
}

func (j jazCash) pay(amount float32) {
	fmt.Println("amount pay with JazzCash : ", amount)
}

type nayaPay struct {
}

func (n nayaPay) pay(amount float32) {
	fmt.Println("amount pay with nayaPay : ", amount)
}
func main() {
	newPay := payment{
		gatWay: nayaPay{},
	}
	newPay.pay(324)
}
