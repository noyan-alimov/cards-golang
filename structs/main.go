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

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateFirstName(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "anderson@mail.com",
			zipCode: 01000,
		},
	}

	alexPointer := &alex
	alexPointer.updateFirstName("John")
	alex.print()
}
