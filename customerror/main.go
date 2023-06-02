package main

import "fmt"

type userNotFound struct {
	name string
}

func (u *userNotFound) Error() string {
	return fmt.Sprintf("%s is not found.\n", u.name)
}

var (
	test = func() error { return &userNotFound{name: "AAA"} }
)

func main() {
	if err := test(); err != nil {
		fmt.Println(err)
	}
}
