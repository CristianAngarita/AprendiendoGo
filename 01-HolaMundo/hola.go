package main

import (
	"fmt"
)

func main() {

	//Atributos
	var name string
	var age int

	//Lee datos desde consola
	fmt.Print("ingrese nombre")
	fmt.Scanln(&name)
	fmt.Print("ingrese edad")
	fmt.Scanln(&age)

	//Imprime en consola.
	fmt.Printf("hola, soy %s y tengo %d años.\n", name, age)

	//Construye el mensaje de saludo completo y lo guarda en la variable 'greeting'
	greeting := fmt.Sprintf("hola, soy %s y tengo %d años.\n", name, age)
	fmt.Println(greeting)

}
