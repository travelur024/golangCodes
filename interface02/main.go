package main

import "fmt"

type User interface {
	Permissons() int //1-5
	Name() string
}

type Admin struct {
	name string
}

type Productor struct {
	name1 string
}

func (this Productor) Permissons() int {
	return 2
}

func (this Admin) Permissons() int {
	return 5
}
func (this Admin) Name() string {
	return this.name
}
func (this Productor) Name() string {
	return this.name1
}

func auth(user User) string {
	if user.Permissons() > 4 {
		return user.Name() + " Tiene permisos de administrador"
	}
	return user.Name() + " No tiene permisos"
}

func main() {
	// admin := Admin{"Jonathan"}
	// fmt.Println(admin.Name())
	// fmt.Println(admin.Permissons())
	// fmt.Println(auth(admin))
	// producer := Productor{"Uriel"}
	// fmt.Println(producer.Name())
	// fmt.Println(producer.Permissons())
	// fmt.Println(auth(producer))

	usuarios := []User{Admin{"Jonathan"}, Productor{"Uriel"}}
	usuarios = append(usuarios, Admin{"Marcos"})

	for i, elem := range usuarios {
		fmt.Println(auth((elem)))
		fmt.Println(usuarios[i].Name())
	}

}
