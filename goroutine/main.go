package main

import (
	"fmt"
	"sync"
)

func hello(group *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, "hello")
	}
	group.Done()
}

func world() {
	for i := 0; i < 5; i++ {
		fmt.Println(i, "world")
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go hello(&wg)
	world()
	wg.Wait()
}
