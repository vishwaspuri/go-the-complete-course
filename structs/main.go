package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// Ways of creating struct instaces
	alex := person{"Alex", "Anderson"}
	vishwas := person{firstName: "Vishwas", lastName: "Puri"}
	fmt.Println(alex, vishwas)
}
