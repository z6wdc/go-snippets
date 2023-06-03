package main

import "fmt"

type User struct {
	name  string
	level int
}

func (u User) String() string {
	return fmt.Sprintf("%s: lv %d\n", u.name, u.level)
}

func main() {
	fmt.Println(User{name: "AAA", level: 25})
}
