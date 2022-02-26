package main

import (
	"fmt"
)

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarAlista(tl *task) {
	t.tasks = append(t.tasks, tl)
}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *taskList) imprimirLista() {
	for _, elem := range t.tasks {
		fmt.Println("Nombre:", elem.nombre)
		fmt.Println("Descripcion:", elem.descripcion)
	}
}

func (t *taskList) imprimirListaCompletados() {

	for _, elem := range t.tasks {
		if elem.completado == true {
			fmt.Println("Nombre:", elem.nombre)
			fmt.Println("Descripcion:", elem.descripcion)
		}
	}
}
func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func (t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func main() {
	t1 := &task{
		nombre:      "Completar mi curso de Go",
		descripcion: "Completar mi curso de Go en una semana",
		completado:  false,
	}
	t2 := &task{
		nombre:      "Completar mi curso de Python",
		descripcion: "Completar mi curso de Python en una semana",
		completado:  false,
	}
	t3 := &task{
		nombre:      "Completar mi curso de NodeJS",
		descripcion: "Completar mi curso de NodeJS en una semana",
		completado:  false,
	}

	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}
	lista.agregarAlista(t3)
	lista.imprimirLista()
	lista.tasks[0].marcarCompleta()
	fmt.Println("Tares completadas")
	lista.imprimirListaCompletados()

}
