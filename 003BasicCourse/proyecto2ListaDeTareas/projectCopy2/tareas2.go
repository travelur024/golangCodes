package main

import "fmt"

//Arreglo de slices con structs para lsitar tareas...

//Struct de la tarea
type task struct {
	name        string
	description string
	complete    bool
}

//Struct de la lista
type taskLista struct {
	tasks []*task
}

//Actualizar estado de tarea
func (t *task) marcarCompleta() {
	t.complete = true
}

//Actualizar descripcion
func (t *task) actualizarDescripcion(descripcion string) {
	t.description = descripcion
}

//Actulaizar nombre
func (t *task) actualizarNombre(nombre string) {
	t.name = nombre
}

//Agregar a la lista
func (t *taskLista) agregarListas(tsk *task) {
	t.tasks = append(t.tasks, tsk)
}

//Eliminar de la lista
func (t *taskLista) eliminarDeListas(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func main() {
	//Tareas
	t1 := &task{
		name:        "Instalar GO in W10",
		description: "Lo antes posible from yesterday1",
		complete:    false,
	}
	t2 := &task{
		name:        "Instalar JS in W10",
		description: "Lo antes posible from yesterday2",
		complete:    false,
	}
	t3 := &task{
		name:        "Instalar PY in W10",
		description: "Lo antes posible from yesterday3",
		complete:    false,
	}
	t4 := &task{
		name:        "Instalar HTML in W10",
		description: "Lo antes posible from yesterday4",
		complete:    false,
	}

	//Lista de tareas
	lista := &taskLista{
		tasks: []*task{
			t1, t2, t3, t4,
		},
	}
	//Funcion para eliminar tareas respecto al indice
	lista.eliminarDeListas(2)
	//Funcion para agregar tareas...
	lista.agregarListas(t1)

	//fmt.Println(t1)
	fmt.Println(lista)

	for index, elem := range lista.tasks {
		fmt.Println("Index:", index, "Elementos:", elem)
	}

}
