package main

import "fmt"

func main() {
	//DEFER (keyword ejecuta la ultima funcion
	//antes de que muera el codigo)
	//Usado para cerrar archivos o conexiones a bases de datos.
	defer fmt.Println("Hola")
	fmt.Println("World")

	//Continue y break
	for i := 0; i < 10; i++ {

		//Continue
		//La instrucción continue se usa cuando se busca omitir la
		//parte restante del bucle, volver a la parte superior de este y
		//continuar con una nueva iteración.
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}
		fmt.Println(i)

		//Break (cuando quieres que no se ejecute el codigo)
		if i == 8 {
			fmt.Println("break")
			break
		}
	}

}
