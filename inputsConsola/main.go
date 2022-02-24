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

	valores := strings.Split(operation, ("+")) //split de texto (+) a partir del simbolo...
	fmt.Println(valores)                       //se muestra en partes
	fmt.Println(valores[0] + valores[1])       // concatena en 2 partes
	operador1, _ := strconv.Atoi(valores[0])   // converitr a entero
	operador2, _ := strconv.Atoi(valores[1])   // converitr a entero
	fmt.Println(operador1 + operador2)         //show resultado aritmetico

}
