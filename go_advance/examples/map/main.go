package main

import "fmt"

type MapStruct struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]int)
	m["1"] = 1
	fmt.Println("map:", m)
	add(m)
	fmt.Println("map:", m)
	n := map[string]int{"1": 1, "2": 2}
	fmt.Println("map:", n)
	add(n)
	fmt.Println("map:", n)
	/*
		map: map[1:1]
		add map: map[1:1 22:22]
		map: map[1:1 22:22]
		map: map[1:1 2:2]
		add map: map[1:1 2:2 22:22]
		map: map[1:1 2:2 22:22]
	*/
	mapStruct := make(map[MapStruct]int)
	a := MapStruct{
		Name: "a",
		Age:  0,
	}
	b := MapStruct{
		Name: "a",
		Age:  1,
	}
	c := MapStruct{
		Name: "a",
		Age:  2,
	}
	mapStruct[a] = a.Age
	mapStruct[b] = b.Age
	mapStruct[c] = c.Age
	fmt.Println("mapStruct:", mapStruct)
}
func change(ma map[string]int) {
	ma["1"] = 22
	fmt.Println("change map:", ma)
}
func add(ma map[string]int) {
	ma["22"] = 22
	fmt.Println("add map:", ma)
}
