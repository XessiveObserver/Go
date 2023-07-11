package main

import (
	"fmt"
	"strings"
)

func main() {

	var name = [3]string{"okello", "joan", "peter"}

	fmt.Println(strings.ToUpper(name[0]))
	fmt.Println(strings.ToTitle(name[1]))
	fmt.Println(strings.ToLower(name[2]))

}
