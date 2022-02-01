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
	aaa := person{
		firstName: "Player",
		lastName:  "aaa",
		contact: contactInfo{
			email:   "aaa@mail.com",
			zipCode: 94000,
		},
	}
	fmt.Println(aaa)

	var bbb person
	fmt.Println(bbb)
	fmt.Printf("%+v\n", bbb)

	// Updating Struct Values
	bbb.firstName = "bbb"
	bbb.lastName = "player"
	bbb.contact = contactInfo{
		email:   "bbb@mail.com",
		zipCode: 12345,
	}
	bbb.print()

	bbbPointer := &bbb

	bbbPointer.updateName("ccc")
	bbb.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}