// Pointers

package main

import "fmt"

func main() {
	age := 20
	his_age := &age
	fmt.Println("The pointer of age is", his_age, "\nand the pointer value is", *his_age)
}
