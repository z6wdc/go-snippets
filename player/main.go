package main

import "fmt"

type player struct {
	name  string
	score int
}

func main() {
	a := player{"aaa", 0}
	fmt.Println(a)

	pointerA := &a
	pointerA.updateScore(1)
	fmt.Println(a)
}

func (p *player) updateScore(s int) {
	(*p).score = s
}
