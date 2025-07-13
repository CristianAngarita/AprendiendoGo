package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

//Opera sobre el puntero de persona
func (p *Persona) diHola() {
	fmt.Println("Hola, mi nombre es", p.nombre)
}

func main() {
	/* Definimos una variable x que luego se va a modificar por medio de una funcion
	y un puntero */
	var x int = 10
	fmt.Println(x)
	editar(&x)
	fmt.Println(x)

	//Instanciamos personas y llamamos el metodo diHola que imprimira un mensaje por un
	//puntero
	pe := Persona{"Cristian", 26, "c11sga@gmail.com"}
	pe.diHola()
}

func editar(x *int) {
	*x = 20
}
