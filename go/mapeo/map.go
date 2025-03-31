package main

import "fmt"

func main() {
	//var mapeo map[string]int
	mapeo1 := make(map[string]int)
	//mapeo2 := map[string]int{"id": 1, "id2": 2}
	mapeo1["id"] = 6
	mapeo1["dos"] = 10
	mapeo1["tres"] = 20
	mapeo1["cuatro"] = 30

	for i, elements := range mapeo1 {
		fmt.Println(i, elements)
	}

}
