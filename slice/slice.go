package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// empty integer slice of size (length) 3
	numbers := make([]int, 3)
	times := 0

	for {
		var command string

		fmt.Println("Please input an integer number.")
		fmt.Println("Or press X to exit.")

		fmt.Scanln(&command)

		if number, err := strconv.Atoi(command); err == nil {
			times++
			if times <= 3 {
				numbers[times-1] = number
			} else {
				numbers = append(numbers, number)
			}

			sortNumbers := make([]int, len(numbers))
			copy(sortNumbers, numbers)
			// Sort the copied slice, not the origin one.
			sort.Slice(sortNumbers, func(i, j int) bool {
				return sortNumbers[i] < sortNumbers[j]
			})

			fmt.Println("The numbers you have input so far")
			fmt.Printf("numbers: %v\n", sortNumbers)
		} else if command == "X" {
			break
		}
	}
}
