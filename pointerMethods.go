// pointer methods

package main

import "fmt"

type Lang struct {
	Name     string
	Creator  string
	YearMade int
}

func (language *Lang) GiveDetails() {
	fmt.Println(language.Name, "was created by", language.Creator, "in", language.YearMade)
}

func main() {
	lang := Lang{"Python", "Guiddo Van Rosum", 1992}
	lang.GiveDetails()
}
