package main

import "fmt"

func double(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + v)
	default:
		fmt.Printf("DO NOT DOUBLE WITH TYPE %T\n", v)
	}
}

func main() {
	double(2)
	double("TEST")
	double(true)
}
