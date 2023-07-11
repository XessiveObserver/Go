package main

import (
	"fmt"
	"strings"
)

type Profile struct {
	Age       int
	Name      string
	Residence string
}

var person = Profile{24, "paul", "area 40"}

func main() {
	fmt.Println(person.Age)
	fmt.Println(strings.Title(person.Name))
	fmt.Println(strings.ToUpper(person.Residence))
}
