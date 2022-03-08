package main

import "fmt"

//Inteface
type Area interface {
	CalcularArea() int
	CalcularPerimetro() int
}

// Cuadrado implementa Area de manera implicita
type Cuadrado struct {
	lado int
}

//Cuadrado area, perimetro
func (c *Cuadrado) CalcularArea() int {
	return c.lado * c.lado
}

func (c *Cuadrado) CalcularPerimetro() int {
	return c.lado * 4
}

// Rectangulo implementa Area de manera implicita y perimetro
type Rectangulo struct {
	lado, base int
}

//Rectangulo area
func (r *Rectangulo) CalcularArea() int {
	return r.lado * r.base
}

func (r *Rectangulo) CalcularPerimetro() int {
	return (r.base * 2) + (r.lado * 2)
}

// Triangulo implementa Area de manera implicita
type Triangulo struct {
	base, altura int
}

//Triangulo area
func (t *Triangulo) CalcularArea() int {
	return (t.base * t.altura) / 2
}
func (r *Triangulo) CalcularPerimetro() int {
	return (r.base * r.altura) % 2
}

func main() {

	figuras := make([]Area, 0)
	figuras = append(figuras, &Cuadrado{lado: 4})
	figuras = append(figuras, &Rectangulo{lado: 4, base: 6})
	figuras = append(figuras, &Triangulo{base: 4, altura: 10})

	// Calcular areas
	for _, f := range figuras {
		fmt.Println("Area:", f.CalcularArea())
		fmt.Println("Perimetro:", f.CalcularPerimetro())
	}

}
