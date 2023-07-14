package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(50 * time.Millisecond)
	end := time.After(1 * time.Second)

out:
	for {
		select {
		case t := <-tick:
			fmt.Println("tick", t)
		case e := <-end:
			fmt.Println("over", e)
			break out
		default:
			fmt.Println(".....")
			time.Sleep(100 * time.Millisecond)
		}
	}
	fmt.Println("##########")
}
