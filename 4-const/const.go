package main

import "fmt"
var names = "sdfsdf"
// helo:="sdfsd" //we can'nt define like thiss
func main(){
	const name = "usama"
	fmt.Println(name)
	const (
		age = 22
		port= 5050
		host = "localhost" //we can group all const 
	) 
	fmt.Println(age)
	fmt.Println(port)
	fmt.Println(host)
}