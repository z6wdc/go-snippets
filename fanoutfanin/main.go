package main

import "fmt"

func producer(first chan<- int, number int) {
	defer close(first)
	for i := 0; i < number; i++ {
		first <- i
	}
}

func plus(first <-chan int, second chan<- int, number int) {
	defer close(second)
	for i := range first {
		second <- i + number
	}
}

func multiple(second <-chan int, third chan<- int, number int) {
	defer close(third)
	for i := range second {
		third <- i * number
	}
}

func main() {
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go producer(first, 10)
	go plus(first, second, 5)
	go multiple(second, third, 3)

	for i := range third {
		fmt.Println(i)
	}
}
