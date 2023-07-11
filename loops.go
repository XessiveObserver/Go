package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {

		if i%2 == 0 {
			fmt.Println(i, "is an even number")
		} else if i%2 != 0 {
			fmt.Println(i, "is an odd number")
		} else {
			fmt.Println(i)
		}
	}
}
