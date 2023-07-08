package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var list []int = make([]int, 0, 3)
	var input string
	for true {
		fmt.Scan(&input)
		if input == "X" {
			break
		}
		var input2 int
		var err error
		input2, err = strconv.Atoi(input)
		if err != nil {
			println("Wrong input")
		}
		list = append(list, input2)
		sort.Ints(list)
		fmt.Println(list)
	}
}
