package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Repasar os.Stdin libreria
	scanner := bufio.NewScanner(os.Stdin) //Andir texto desde la terminal
	//scanner := bufio.NewScanner(strings.NewReader("Hola khe hace")) // anadir texto desde un string
	scanner.Scan()              //Espera a ingresar en terminal y escanear...
	operation := scanner.Text() //pase de parametros string
	//fmt.Println(scanner.Bytes()) //impresion del scaner en bytes
	fmt.Println(operation) //impresion del escaner

	operador := "p"                               //Asigna operacion
	valores := strings.Split(operation, operador) //split de texto (+) a partir del simbolo...
	fmt.Println(valores)                          //se muestra en partes
	fmt.Println(valores[0] + valores[1])          // concatena en 2 partes
	operador1, err1 := strconv.Atoi(valores[0])   // converitr a entero
	//Error1, ya que strconv.Atio devuelve 2 valores entero y error
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println("Operador 1 numero:", operador1)
	}

	operador2, err2 := strconv.Atoi(valores[1]) // converitr a entero
	//Error2, ya que strconv.Atio devuelve 2 valores entero y error
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Operador 2 numero:", operador2)
	}

	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
	case "-":
		fmt.Println(operador1 - operador2)
	case "/":
		fmt.Println(operador1 / operador2)
	case "*":
		fmt.Println(operador1 * operador2)
	default:
		fmt.Println(operador, "No estas soportado")

	}

}
