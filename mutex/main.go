package main

import (
	"fmt"
	"sync"
)

const KEY = "key"

type counter struct {
	record map[string]int
	mux    sync.Mutex
}

func (c *counter) plus(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.record[key]++
}

func (c *counter) printValueBy(key string) {
	fmt.Println(c.record[key])
}

func main() {
	var wg sync.WaitGroup
	c := counter{record: make(map[string]int)}
	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			c.plus(KEY)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c.plus(KEY)
		}
		wg.Done()
	}()
	wg.Wait()
	c.printValueBy(KEY)
}
