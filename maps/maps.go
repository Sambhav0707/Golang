package main

import "fmt"

func main() {

	map1 := map[string]int{}

	map1["apple"] = 1
	map1["banana"] = 2
	map1["cherry"] = 3

	fmt.Println(map1)
	for key, val := range map1 {
		fmt.Println(key, val)
	}
}
