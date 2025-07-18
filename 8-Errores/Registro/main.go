package main

import "log"

func main() {
	log.Fatal("AA")                     // Imprime el mensaje y TERMINA el programa
	log.Println("Otro msj de registro") // Nunca se ejecuta porque ya terminó con la linea anterior.
}
