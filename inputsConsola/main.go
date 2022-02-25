package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Struc
type calc struct {
	sum string
}

func (calc) operate(entrada string, operador string) int {
	entradaLimpiea := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpiea[0])
	operador2 := parsear(entradaLimpiea[1])
	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
		return operador1 + operador2
	case "-":
		fmt.Println(operador1 - operador2)
		return operador1 - operador2
	case "/":
		fmt.Println(operador1 / operador2)
		return operador1 / operador2
	case "*":
		fmt.Println(operador1 * operador2)
		return operador1 * operador2
	default:
		fmt.Println(operador, "No estas soportado")
		return 0
	}

}

func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada) // converitr a entero
	return operador
}

func leerEntrada() string {
	//Repasar os.Stdin libreria
	scanner := bufio.NewScanner(os.Stdin) //Andir texto desde la terminal
	//scanner := bufio.NewScanner(strings.NewReader("Hola khe hace")) // anadir texto desde un string
	scanner.Scan() //Espera a ingresar en terminal y escanear...
	return scanner.Text()
}

func main() {

	entrada := leerEntrada()
	operador := leerEntrada()
	fmt.Println(entrada)
	fmt.Println(operador)
	c := calc{}
	c1 := c.operate(entrada, operador)
	println(c1)
}
