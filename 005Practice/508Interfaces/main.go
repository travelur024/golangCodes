package main

import (
	"fmt"
	"strconv"
)

//Interface
type ReadDescription interface {
	Reading() string
}

type Book struct {
	Pages int
	Name  string
	Type  string
}

type Article struct {
	Name string
	Type string
}

//Metodos Objecto book
func (b Book) Reading() string {
	return "Name: " + b.Name + " Type: " + b.Type + " Counts with:  " + strconv.Itoa(b.Pages) + " Pages"
}

//Metodos Objecto articule
func (a Article) Reading() string {
	return "Name: " + a.Name + " Type: " + a.Type
}

//Funcion que implementa Read...
func Print(a ReadDescription) {
	fmt.Println(a.Reading())
}

func main() {
	// Book001 := Book{
	// 	Pages: 269,
	// 	Name:  "El unico",
	// 	Type:  "Romantico",
	// }

	// Print(Book001)

	sliceBooks := []Book{{Pages: 269,
		Name: "El unico",
		Type: "Romantico"}, {Pages: 365,
		Name: "El indiferente",
		Type: "Drama"}}

	for i, book := range sliceBooks {
		fmt.Println("Book:", i+1)
		Print(book)
	}

	sliceArticles := []Article{
		{Name: "Scientific Topics", Type: "Science"},
		{Name: "Non Scientific Topics", Type: "Not Science"},
	}

	for i, article := range sliceArticles {
		fmt.Println("Article:", i+1)
		Print(article)
	}

}
