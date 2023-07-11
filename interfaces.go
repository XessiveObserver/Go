// interfaces

package main

import "fmt"

type Speaker interface {
	Speak()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Speak() {
	fmt.Println("Your", p.Name, "of", p.Age, "years old!")
}

func SaySomething(s Speaker) {
	s.Speak()
}

func main() {
	xessive := Person{"Observer", 33}
	SaySomething(xessive)
}
