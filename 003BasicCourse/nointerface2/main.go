package main

import "fmt"

//Animales que indican como se estan moviendo - Ejemplo interfaces
//Structs basicos de animales
type animal interface {
	//Todo struct que use este metodo se considera animal
	mover() string
}

type perro struct {
}

type pez struct {
}

type pajaro struct {
}

//Metodos de animales- Con la accion que necesitamos
func (perro) mover() string {
	return "Soy un perro y camino"
}
func (pez) mover() string {
	return "Soy un pez y nado"
}
func (pajaro) mover() string {
	return "Soy un pajaro y vuelo"
}

//Metodo que imprima que esta ocurriendo como se mueve el animal
func moverAnimal(a animal) {
	fmt.Println(a.mover())
}

func main() {
	p := perro{}
	moverAnimal(p)
	pe := pez{}
	moverAnimal(pe)
	pa := pajaro{}
	moverAnimal(pa)
	fmt.Println("\n")

	//Forma con maps de ingresar
	animalitos := make([]animal, 0)
	animalitos = append(animalitos, p, pe, pa, pa, p)
	for i, _ := range animalitos {

		moverAnimal(animalitos[i])
	}

	//Just for practice...
	// circo := make(map[string]*animal, 0)
	// circo["Jonathan"] = &animalitos[0]
	// fmt.Println("\n")
	// moverAnimal(*circo["Jonathan"])
	// fmt.Println(&circo)

}
