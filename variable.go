package main

import "fmt"

var quantity = 4
var length, width = 1.2, 2.4
var customerName = "Damon Cole"

func main() {
	fmt.Println(customerName)
	fmt.Println("Has ordered", quantity, "sheets")
	fmt.Println("Each with an area of")
	fmt.Println(length*width, "Square meters")
}
