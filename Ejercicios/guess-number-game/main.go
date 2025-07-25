package main

import (
	"fmt"
	"math/rand"
)

func main() {
	play()
}

func play() {
	num := rand.Intn(100)
	var numIngresado int
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Ingresa un número (intentos restantes: %d):", maxIntentos-intentos+1)
		fmt.Scanln((&numIngresado))

		if numIngresado == num {
			fmt.Println("Felicidades, adivinaste elnumero")
			jugarNuevamente()
			return
		} else if numIngresado < num {
			fmt.Println("El numero a adivinar es mayor.")
		} else if numIngresado > num {
			fmt.Println("El numero a adivinar es menor.")
		}
	}
	println("Se acabaron los intentos. El numero era: ", num)
	jugarNuevamente()
}

func jugarNuevamente() {
	var eleccion string
	fmt.Print("Quiere jugar otra vez (s/n)")
	fmt.Scanln(&eleccion)

	switch eleccion {
	case "s":
		play()
	case "n":
		fmt.Println("Gracias por jugar")
	default:
		fmt.Println("Eleccion invalida, Intentar nuevamente")
		jugarNuevamente()
	}
}
