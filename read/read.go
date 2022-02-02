package main

import (
	"fmt"
	"os"
	"strings"
)

type nameStruct struct {
	fname string
	lname string
}

func main() {
	var file string
	var filename string
	var names []nameStruct

	for {
		fmt.Println("Please input the file name you want to read.")
		fmt.Println("Or enter X to exit.")
		fmt.Scanln(&filename)

		if filename == "X" {
			return
		}

		if data, err := os.ReadFile(filename); err == nil {
			file = string(data)
			break
		} else {
			fmt.Println(err)
		}
	}

	splitByLine := strings.Split(file, "\n")
	for _, name := range splitByLine {
		splitBySpace := strings.Split(name, " ")
		names = append(names,
			nameStruct{fname: splitBySpace[0], lname: splitBySpace[1]})
	}

	for _, name := range names {
		fmt.Printf("%+v\n", name)
	}
}
