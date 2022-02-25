package main

import (
	"fmt"
)

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t task) marcarCompleta() {
	t.completado = true
}

func (t task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t := task{
		nombre:      "Completar mi curso de Go",
		descripcion: "Completar mi curso de Go en una semana",
	}
	fmt.Printf("%+v\n", t)
	//fmt.Println(reflect.TypeOf(t.nombre))
	t.marcarCompleta()
	t.actualizarNombre("Finalizar curso de go")
	//t.nombre = "Finalizar curso de go"
	t.actualizarDescripcion("Completar curso cuanto antesbb")
	fmt.Printf("%+v\n", t)
}
