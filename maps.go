package main

import "fmt"

func main() {
	profile := make(map[string]int)

	profile["fear"] = 1000
	profile["hate"] = 25

	fmt.Println(profile)
	fmt.Println(profile["fear"])

	status := make(map[int]string)

	status[24] = "age"
	status[35] = "married"

	fmt.Println(status)
	fmt.Println(status[35])
}
