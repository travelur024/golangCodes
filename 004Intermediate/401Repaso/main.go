package main

import (
	"fmt"
	"strconv"
)

func main() {
	const z int = 3
	var x int
	x = 8
	y := 7
	fmt.Println(x, z)
	fmt.Println(y)
	myValue, err := strconv.ParseInt("889s", 0, 64)
	if err != nil {
		fmt.Println("Ponte trucha: ", err)
	} else {
		fmt.Println(myValue)
	}
	//Map
	m := make(map[string]int)
	m["Key"] = 6
	fmt.Println(m["Key"])
	//Slice
	s := []int{9, 8, 7, 6}
	for index, value := range s {
		fmt.Println("\n", index, value)
	}
	s = append(s, 16)
	s = append(s, 17)
	s = append(s, 18)
	for index, value := range s {
		fmt.Println("\n", index, value)
	}
}
