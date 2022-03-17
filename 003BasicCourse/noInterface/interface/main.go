package main

import "fmt"

//Animales que indican como se estan moviendo - Ejemplo interfaces
//Aplicando la interface con meotodos independiente...

type Animal interface {
	Accion() string
}
type perro struct {
}

type pez struct {
}

type pajaro struct {
}

//Metodos de animales- Con la accion que necesitamos
func (perro) Accion() string {
	return "Hola soy un perro Gua! Gua!"
}

func (pez) Accion() string {
	return "Hola soy un pez Agua! Agua!"
}
func (pajaro) Accion() string {
	return "Hola soy un pajaro Pico! Pico!"
}

func moverAnimal(a Animal) {
	fmt.Println(a.Accion())
}

func main() {
	animales := make([]Animal, 0)
	animales = append(animales, &perro{})
	animales = append(animales, &pez{})
	animales = append(animales, &pajaro{})

	// for i, elem := range animales {
	// 	fmt.Println(animales[i], elem)
	// }
	moverAnimal(animales[2])
	moverAnimal(animales[1])
	moverAnimal(animales[0])

}
