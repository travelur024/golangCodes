package main

import (
	"fmt"
	"regexp"
	"strings"
)

func palFunc(text string) {
	//Test of replacement 2
	text2 := strings.ReplaceAll(text, " ", "") // replacement (more easy)
	text2 = strings.ToLower(text2)
	fmt.Println(text2) // text2 prueba tested
	//
	var textInvert string
	r := regexp.MustCompile(" ")            //Search a caracter in a text
	newText := r.ReplaceAllString(text, "") // replace with the character search
	newText = strings.ToLower(newText)
	for i := len(newText) - 1; i >= 0; i-- {
		textInvert += string(newText[i])

	}
	if newText == textInvert {
		fmt.Println("Es un palindromo", newText)
	} else {
		fmt.Println("No es un palindromo", newText)
	}

}

func main() {
	//Error.../ No lee sin espacios...
	fmt.Println("Escribe una frase:")
	var text01 string
	fmt.Scanf("%s", &text01)
	palFunc(text01) //Call function and send text

	fmt.Println(text01)
}
