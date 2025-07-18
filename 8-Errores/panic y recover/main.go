package main

import "fmt"

/* panic: Detiene la ejecución normal del programa con un mensaje de error grave.

recover: Permite atrapar un panic y evitar que el programa se cierre.

defer: Retrasa la ejecución de una función hasta que termina la función actual. */

func div(divi, divis int) {
	// Defer que captura cualquier pánico ocurrido en esta función
	defer func() {
		if r := recover(); r != nil { // Si ocurre un panic, se recupera aquí y se imprime el mensaje
			fmt.Println(r)
		}
	}()
	validateZero(divis)
	fmt.Println(divi / divis)
}

// Esta función lanza un panic si el valor es 0
func validateZero(divi int) {
	if divi == 0 {
		panic("No se puede dividir en 0")
	}
}
func main() {
	div(100, 10)
	div(100, 0)
	div(100, 0)
	div(100, 2)
}
