package main

import "fmt"

func main(){
	name := "Nut Jobber"
	age := 25
	length, width := 2.4, 4.8

	fmt.Println("Your name is", name)
	fmt.Println("Your", age, "years old")
	fmt.Println("You live in", length * width, "square meter's house")
}