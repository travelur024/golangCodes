package main

import "fmt"

func main() {
	//Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	//Slice
	slice := []int{9, 8, 7, 6, 5, 4, 3}
	fmt.Println(slice, len(slice), cap(slice))

	//Metodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[3:])

	//Append
	slice = append(slice, 2)
	fmt.Println(slice, len(slice), cap(slice))

	//New slice
	newSlice := []int{1, 0}
	slice = append(slice, newSlice...)
	fmt.Println(slice, len(slice), cap(slice))
}

// var array1 [2]float32
