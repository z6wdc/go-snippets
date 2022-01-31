package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Please input a string.")
	r := bufio.NewReader(os.Stdin)
	text, _ := r.ReadString('\n')

	// ReadString will read string including delimiter
	cleanText := strings.TrimSuffix(text, "\n")

	if findian(cleanText) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

func findian(text string) bool {
	lowerText := strings.ToLower(text)
	return strings.HasPrefix(lowerText, "i") &&
		strings.HasSuffix(lowerText, "n") &&
		strings.Contains(lowerText, "a")
}
