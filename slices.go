package main

import "fmt"

func main() {
	names := []string{"fear", "deceit", "lies"}

	fmt.Println(names[0:2])

	moreNames := append(names, "hate", "deception")

	fmt.Println(moreNames[1:5])
}
