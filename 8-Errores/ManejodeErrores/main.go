package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Error en caso de conversión de una cadea a int.

	str := "123f"
	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("numeros", num)

}
