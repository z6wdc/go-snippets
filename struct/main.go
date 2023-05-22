package main

import "fmt"

type User struct {
	Id, Year int
	Name     string
}

func main() {
	u1 := User{Id: 0, Year: 10}
	fmt.Println(u1)
	fmt.Println(u1.Id, u1.Year)
	u1.Year = 20
	fmt.Println(u1.Id, u1.Year)

	u2 := User{}
	fmt.Printf("%T %v\n", u2, u2)

	var u3 User
	fmt.Printf("%T %v\n", u3, u3)

	u4 := new(User)
	fmt.Printf("%T %v\n", u4, u4)

	u5 := &User{}
	fmt.Printf("%T %v\n", u5, u5)
}
