package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}
	jim.print()
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
}

// Recievers with structs
func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// Pass by reference
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
