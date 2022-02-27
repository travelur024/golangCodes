package main

import (
	"fmt"
	"io"
	"net/http"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

type escritorWeb struct {
}

func (escritorWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return 0, nil
}

func main() {
	respuesta, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
	}
	slice1 := []Writer{}
	slice1 = append(slice1, escritorWeb{})
	io.Copy(slice1[0], respuesta.Body)

}
