package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type name struct {
	firstname string
	lastname  string
}

func main() {

	var input string
	fmt.Println("Enter the name of the file: ")
	fmt.Scan(&input)
	content, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(content), "\n")

	var names []name
	for _, line := range lines {
		var single_name = strings.Split(line, " ")
		if len(single_name) >= 2 {
			names = append(names, name{firstname: single_name[0], lastname: single_name[1]})
		}

	}
	fmt.Print(names, "\n")
}
