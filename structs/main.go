package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	var alex person
	fmt.Println(alex)
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Printf("%+v\n", alex)
}
