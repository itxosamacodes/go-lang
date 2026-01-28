package main

import "fmt"

type animal interface {
	sound(sound string)
}
type dog struct{}

func (d dog) sound(sound string) {
	fmt.Println("dok make : ", sound)
}

type cat struct{}

func (c cat) sound(sound string) {
	fmt.Println("cat make : ", sound)
}

func main() {
	var p animal
	p = cat{}
	p.sound("meow")
}
