package main

import (
	"fmt"
	"os"
)

func test() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("start")
	fmt.Println("finish")
}

func main() {
	test()
	file, _ := os.Open("./defer/main.go")
	defer file.Close()

	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}
