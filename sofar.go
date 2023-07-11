// interfaces
package main

import "fmt"

type Chef interface {
	Cook()
}

type ChefStatus struct {
	Name    string
	Age     int
	favDish string
	Height  float64
}

func (checfDetails ChefStatus) Cook() {
	fmt.Println("Chef", checfDetails.Name)
	fmt.Println(checfDetails.Age, "years of age")
	fmt.Println(checfDetails.Height, "meters of height")
	fmt.Println("Enjoys", checfDetails.favDish)
}

func WhoAreYou(personality Chef) {
	personality.Cook()
}

func main() {
	profession := ChefStatus{"Egudu", 55, "Pasted Meat", 5.8}
	WhoAreYou(profession)
}
