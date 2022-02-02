package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	data := make(map[string]string)

	fmt.Println("Please input name:")
	name := readString()

	fmt.Println("Please input address:")
	address := readString()

	data["name"] = name
	data["address"] = address

	if jsondata, err := json.Marshal(data); err == nil {
		fmt.Println("JSON:", string(jsondata))
	} else {
		fmt.Println(err)
	}
}

func readString() string {
	r := bufio.NewReader(os.Stdin)
	input, _ := r.ReadString('\n')

	// ReadString will read string including delimiter
	return strings.TrimSuffix(input, "\n")
}
