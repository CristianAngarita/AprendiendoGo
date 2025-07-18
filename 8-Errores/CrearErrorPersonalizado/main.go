package main

import (
	"errors"
	"fmt"
)

// manejo de error division por 0
func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		//Se crea un error basico, en caso de que se divida por 0
		return 0, errors.New("No se puede dividir por 0")
	}
	//Resultado + nil todo OK
	return dividendo / divisor, nil
}
func main() {
	retultado, err := divide(10, 0)
	//Si hay error devuelve diferente a nil
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	//Si no hay error devuelve el resultado esperado.
	fmt.Println("Resultado: ", retultado)
}
