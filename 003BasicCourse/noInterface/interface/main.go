package main

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
func (p *perro) Accion() string {
	return "Soy un perro y ladro"
}
func (p *pez) Accion() string {
	return "Soy un pez y nado"
}
func (p *pajaro) Accion() string {
	return "Soy un pajaro y vuelo"
}

func moverAnimal(a Animal) {
	println(a.Accion())
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
