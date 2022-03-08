package main

import "fmt"

func main() {
	Server1 := NewServer(":3000")
	Server1.Handle("/", HandleRoot)
	Server1.Handle("/home", HandleHome)
	Server1.Listen()

	fmt.Println("Aver al cine")
}
