package main

import (
	"fmt"
	"library/animal"
	"library/book"
)

func main() {
	//Creamos nuevos libros
	myBook := book.NewBook("Te deseo el amor", "Papa Francisco", 303)
	myBook2 := book.NewBook("Constructor", "Simualado", 2)
	myBook.PrintInfo()

	println("========================================================")

	myBook2.PrintInfo()
	myBook2.SetTitle("Constuctor 2")
	fmt.Println(myBook2.GetTitle())

	println("========================================================")

	myBook2.PrintInfo()

	println("======================Composición==================================")
	// Este constructor utiliza composición: TextBook contiene un Book internamente.
	myTextBook := book.NewTextBook("Charlie Y La Fabrica De Chocolate", "Roald Dahl", 231,
		"richmond", "Infantil")
	myTextBook.PrintInfo()

	println("==========================Polimorfismo==============================")
	book.Print(myBook)
	fmt.Println()
	book.Print(myTextBook)

	println("==========================Interface==============================")
	//Interfaces: Animal requiere que se implemente el método Sonido().
	miperro := animal.Perro{Nombre: "firulais"}
	migato := animal.Gato{Nombre: "cocoa"}

	animal.HacerSonido(&miperro)
	animal.HacerSonido(&migato)

	println("===========================Mas objetos + iterar============================")

	animales := []animal.Animal{
		&animal.Perro{Nombre: "Luna"},
		&animal.Gato{Nombre: "Bigotes"},
		&animal.Gato{Nombre: "Nikol"},
		&animal.Perro{Nombre: "Goku"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}
}
