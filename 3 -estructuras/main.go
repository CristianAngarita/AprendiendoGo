package main

import "fmt"

//Definimos una estructura para persona
type Persona struct {
	nombre string
	edad   int
	correo string
}

func main() {
	/*var p Persona
	p.nombre = "Cristian"
	p.edad = 2
	p.correo = "angarita@gmail.com"*/

	//Definimos una instacia el cual mostramos y modificamos
	p := Persona{"Cristian", 26, "angar@gmail.com"}
	fmt.Println(p)
	fmt.Println(p.nombre)
	p.edad = 30
	fmt.Println(p)

	//Definimos otra instancia
	p2 := Persona{"Fredy", 21, "garitax@gmail.com"}
	fmt.Println(p2)
}
