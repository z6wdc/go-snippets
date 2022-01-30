package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	fmt.Println("Please input a string.")
	r := bufio.NewReader(os.Stdin)
	text, _ := r.ReadString('\n')

	// ReadString will read string including delimiter
	cleanText := strings.TrimSuffix(text, "\n") 
	lowerText := strings.ToLower(cleanText)

	var isFound bool
	isFound = strings.HasPrefix(lowerText, "i") &&
		strings.HasSuffix(lowerText, "n") &&
		strings.Contains(lowerText, "a")

	if isFound {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
