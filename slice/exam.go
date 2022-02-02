package main

import "fmt"

func main() {
	exam1()

	exam2()

	exam3()
}

func exam1() {
	x := [...]int{4, 8, 5}
	y := x[0:2]
	z := x[1:3]
	y[0] = 1
	z[1] = 3
	fmt.Print(x)
}

func exam2() {
	x := [...]int{1, 2, 3, 4, 5}
	y := x[0:2]
	z := x[1:4]
	fmt.Print(len(y), cap(y), len(z), cap(z))
}

func exam3() {
	s := make([]int, 0, 3)
	s = append(s, 100)
	fmt.Println(len(s), cap(s))
	fmt.Println(s)
}
