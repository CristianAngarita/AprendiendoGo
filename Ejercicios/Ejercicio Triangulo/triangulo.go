package main

import (
	"fmt"
	"math"
)

func main() {
	const presicion = 2
	var lado1 float32
	var lado2 float32
	var area float32
	var hipo float32
	var perim float32

	fmt.Print("ingrese lado")
	fmt.Scanln(&lado1)
	fmt.Print("ingrese lado")
	fmt.Scanln(&lado2)

	hipo = float32(math.Sqrt(math.Pow(float64(lado1), 2) + math.Pow(float64(lado2), 2)))
	area = (lado1 * lado2) / 2
	perim = lado1 + lado2 + hipo

	fmt.Printf("Área: %.*f\nPerimetro: %.*f", presicion, area, presicion, perim)

}
