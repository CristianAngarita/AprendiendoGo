package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Funcion para imprimir los valores pasados en el argumente,
// ...interface permite obtener un cantidad variable de argumentos de culalquier tipo
func PrintList(list ...interface{}) {
	for _, value := range list {
		fmt.Println(value)
	}
}

// ...any funciona igual que el ...interface solo que se implementa desde Go 1.18+
func PrintListAny(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
}

// funcion generica con un parametro detipo T, el cual puede ser int o float
// nums ...T Es un parámetro variádico: puedes pasar uno o más valores de tipo T
func Sum[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, nun := range nums {
		total += nun
	}
	return total
}

type integer int

// Estás creando una interfaz de restricción de tipos para usar con generics en Go.
type Numbers interface {
	~int | ~float64 | ~float32 | ~uint
}

// Includes verifica si un valor está presente dentro de un slice.
// La función es genérica y funciona con cualquier tipo que sea comparable (como int, string, float, etc.).
func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

// Filter recibe un slice de elementos y una función callback.
// Devuelve un nuevo slice con los elementos que cumplan una condición definida en el callback
func Filter[T constraints.Ordered](list []T, callback func(T) bool) []T {
	// Creamos un slice vacío llamado result, del mismo tipo T, con capacidad inicial igual a list.
	result := make([]T, 0, len(list))

	for _, item := range list {

		if callback(item) {
			// Si el callback devuelve true, agregamos el item al slice res
			result = append(result, item)
		}
	}
	return result
}

type Product[T uint | string] struct {
	Id    T
	Desc  string
	Price float32
}

func main() {
	/* PrintList("Cristian", "Angarita")
	PrintList(1, 3, 5, 7)
	PrintList(1, 3, 5, 7, "Steven")

	fmt.Println("*****************")
	PrintListAny("Cristian", "Angarita")
	PrintListAny(1, 3, 5, 7)
	PrintListAny(1, 3, 5, 7, "Steven")

	var num1 integer = 100
	var num2 integer = 400
	fmt.Println(Sum[int](5, 3, 5, 4))
	fmt.Println(Sum[float64](5, 3, 5, 7.5, 10.2))

	fmt.Println(Sum(num1, num2)) */

	strings := []string{"a", "b", "c", "d"}
	numbers := []int{1, 2, 3, 4, 5, 6}
	/*
		fmt.Println(Includes(strings, "a"))
		fmt.Println(Includes(strings, "z"))

		fmt.Println(Includes(numbers, 1))
		fmt.Println(Includes(numbers, 7)) */
	//func(value int) bool { return value > 3 } esto es una funcion anonima o lambda
	// La función Filter recibe un slice de tipo T y una función callback que recibe un valor de tipo T
	// y retorna un valor booleano que indica si el valor cumple con una condición.

	fmt.Println(Filter(numbers, func(value int) bool {
		// Esta es una función anónima (sin nombre), que recibe un valor entero y retorna
		// verdadero si el valor es mayor que 3. Esta función es pasada como argumento a Filter.
		return value > 3
	}))
	fmt.Println(Filter(strings, func(value string) bool { return value > "b" }))

	producto := Product[uint]{1, "Camisa", 50}
	producto2 := Product[string]{"AAASS-aAAs-WWER", "Camisa", 50}
	fmt.Println(producto, producto2)
}
