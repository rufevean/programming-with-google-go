package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Println("Enter a string: ")
	fmt.Scan(&input)
	input = strings.ToLower(input)
	if input[0] == 'i' && input[len(input)-1] == 'n' && strings.Contains(input, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
