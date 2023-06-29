package main

import "fmt"

func total(number int, c chan int) {
	sum := 0
	for i := 0; i < number; i++ {
		sum += i
		c <- sum
	}
	close(c)
}

func main() {
	ch := make(chan int)
	go total(5, ch)
	for i := range ch {
		fmt.Println(i)
	}
}
