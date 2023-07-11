// Functions
package main

import (
	"fmt"
	"strings"
)

func name(firstName, lastName string) {
	var fullName = firstName + " " + lastName
	fmt.Println(strings.Title(fullName))
}

func add(a, b int) {
	fmt.Println(a + b)
}

func main() {
	name("okello", "francis")
	add(4, 8)
}
