package main

import "fmt"

type Party struct {
	Invitados int
	NameParty string
	Tema      string
	status    string
}

func (p *Party) RunParty() {
	fmt.Printf("La fiesta cuenta con %v invitados \n", p.Invitados)
	fmt.Printf("El nombre sera: %v \n", p.NameParty)
	fmt.Printf("El tema: %v \n", p.Tema)
}

func (p *Party) AsignarParty(i int, n string, t string) {
	p.Invitados = i
	p.NameParty = n
	p.Tema = t
	p.status = "Asigando"
}

func main() {
	Party1 := Party{}
	fmt.Println(Party1)
	Party1.AsignarParty(12, "La alocada", "Comics")
	fmt.Println(Party1)
	Party1.RunParty()
}
