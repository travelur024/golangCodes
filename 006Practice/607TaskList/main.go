package main

import "fmt"

//task
type task struct {
	nombre      string
	descripcion string
	completado  bool
}

//tasklist
type taskList struct {
	tasks []*task
}

//Metodo agregar a la taskList
func (t taskList) agregarLista(tl *task) {
	t.tasks = append(t.tasks, tl)
}

//imprimir lista
func (t *taskList) imprimirLista() {
	for _, elem := range t.tasks {
		fmt.Println("Nombre:", elem.nombre)
		fmt.Println("Descripcion:", elem.descripcion)
	}
}

//Imprimir lista completados
func (t *taskList) imprimirListaCompletados() {
	for _, elem := range t.tasks {
		if elem.completado == true {
			fmt.Println("Nombre:", elem.nombre)
			fmt.Println("Descripcion:", elem.descripcion)
		}
	}
}

//Metodos atajos para cada task
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

//Main

func main() {
	//Tareas creadas...
	t1 := &task{
		nombre:      "Curso Golang",
		descripcion: "Completar curso de Golang",
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

	lista := &taskList{tasks: []*task{t1, t2}}
	lista.agregarLista(t3)
	lista.imprimirLista()
	lista.tasks[0].marcarCompleta()
	lista.tasks[1].marcarCompleta()
	lista.tasks[2].marcarCompleta()

	//Conocer las tareas de cada persona
	mapaTareas := make(map[string]*taskList)

	mapaTareas["Jonathan"] = lista
	t4 := &task{
		nombre:      "Completar mi curso de Java",
		descripcion: "Completar mi curso de Java en una semana",
		completado:  false,
	}
	t5 := &task{
		nombre:      "Completar mi curso de C#",
		descripcion: "Completar mi curso de C# en una semana",
		completado:  false,
	}
	lista2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}

	//Actualizando tareas...
	lista2.tasks[0].marcarCompleta()
	//lista2.tasks[1].marcarCompleta()
	mapaTareas["Ernesto"] = lista2
	//Impresion de tareas indivudiales con maps...
	// fmt.Println("Tareas de Jonathan")
	// mapaTareas["Jonathan"].imprimirLista()
	// fmt.Println("Tareas de Erneso")
	// mapaTareas["Ernesto"].imprimirLista()

	//Impresion del map con For range
	for i, elem := range mapaTareas {
		fmt.Println("Encargado: ", i)
		elem.imprimirListaCompletados()
	}

}
