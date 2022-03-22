package main

import "fmt"

type opera func(int, int) int

func checkSum() opera {
	return func(i1, i2 int) int {
		i3 := i1 + i2
		return i3
	}
}

func main() {
	res := checkSum()
	sum1 := res(2, 3)
	sum2 := res(3, 3)
	sum3 := res(2, 5)

	fmt.Println(sum1, sum2, sum3)
}
