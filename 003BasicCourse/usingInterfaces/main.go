package main

import (
	"fmt"
	"io"
	"net/http"
)

type webWriter struct {
}

func (webWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p)) //string() convierte los datos byte
	return 0, nil
}

func main() {
	//Generando solititud a goolge http.Get
	respuesta, err := http.Get("http://google.com")
	//respuesta contiene los datos que necesitamos leer
	if err != nil {
		fmt.Println(err)
	}
	//Anadiendo metodo Writer
	e := webWriter{}
	io.Copy(e, respuesta.Body)
}

//Handler cosas para gruadar o cmabiar data
//Mineword para borrar o cabmiar datos de la base de datos
