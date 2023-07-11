// methods

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Speake() {
	fmt.Println("Hello from", p.Name)
	fmt.Println("Who's", p.Age, "years old")
}

func main() {
	xessive := Person{"Observer", 33}
	xessive.Speake()
}
