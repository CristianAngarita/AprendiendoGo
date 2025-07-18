package main

import (
	"fmt"
	"os"
)

func main() {
	/*Si tienen todos defer se agregan a una pila de ejecución
	defer fmt.Println(3)
	defer fmt.Println(1)
	defer fmt.Println(2)*/

	// se usa para posponer la ejecución de una función hasta que la funcion que la contiene halla finalizado.
	defer fmt.Println("Final")
	// Se crea (o sobrescribe si ya existe) un archivo llamado "hola.txt"
	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Inicio")

	// Defer pospone el cierre del archivo hasta que main termine
	defer file.Close()

	//Se escribe dentro del archivo
	_, err = file.Write([]byte("Hola, Cris"))
	if err != nil {
		fmt.Println(err)
		return
	}

}
