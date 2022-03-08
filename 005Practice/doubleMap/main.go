package main

import "fmt"

type double struct {
	map2 map[string]map[string]int
}

//Creating a map
func newMap() *double {
	return &double{
		map2: make(map[string]map[string]int),
	}
}

func main() {
	//Mapa doble
	m2 := newMap()
	m3 := make(map[string]map[string]int)

	m3["1"] = make(map[string]int)
	m3["2"] = make(map[string]int)
	m3["1"]["Juan"] = 2
	m3["1"]["Roberto"] = 2
	m3["1"]["Manuel"] = 2
	m3["2"]["Juan2"] = 1
	m3["2"]["Roberto2"] = 1
	m3["2"]["Manuel2"] = 1
	//
	//m2.map2["Group1"]["Juan"] = 32
	fmt.Println(m2, m3)
}
