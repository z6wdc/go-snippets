package main

import (
	"fmt"
)

func main() {
	var foodPrice map[string]int

	fmt.Println(foodPrice)

	foodRate := make(map[string]int)

	fmt.Println(foodRate)

	foodRate["rice"] = 3

	fmt.Println(foodRate)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	fmt.Println(colors)

	delete(colors, "green")

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
