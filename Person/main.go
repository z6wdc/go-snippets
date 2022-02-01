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
	tom := person{
		firstName: "Hank",
		lastName:  "Tom",
		contact: contactInfo{
			email:   "tom@mail.com",
			zipCode: 94000,
		},
	}
	fmt.Println(tom)

	var captain person
	fmt.Println(captain)
	fmt.Printf("%+v\n", captain)

	// Updating Struct Values
	captain.firstName = "keroro"
	captain.lastName = "captain"
	captain.contact = contactInfo{
		email:   "keroro@mail.com",
		zipCode: 12345,
	}
	captain.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}