package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string
	fmt.Scan(&name)
	fmt.Scan(&address)
	map_1 := make(map[string]string)
	map_1["name"] = name
	map_1["address"] = address

	a, _ := json.Marshal(map_1)
	fmt.Println(string(a))

}
