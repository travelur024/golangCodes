package main

import "fmt"

type area func(int, int) int

func main() {
	Rectangulo := getAreaRecta()
	Triangulo := getAreaTrian()
	print(2, 3, Rectangulo)
	print(5, 6, Triangulo)
}

func print(x, y int, a area) {
	fmt.Printf("Area is: %d\n", a(x, y))
}

func getAreaRecta() area {
	return func(b, h int) int {
		return b * h
	}
}

func getAreaTrian() area {
	return func(b, h int) int {
		return (b * h) / 2
	}
}
