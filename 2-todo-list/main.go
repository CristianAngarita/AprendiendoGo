package main

import (
	"fmt"
)

func main() {
	//rebanadas()
	//rebanadasMakeCopy()
	//mapas()
	//matriz()

}

func matriz() {
	//matriz con 5 elementos
	var a = [...]int{10, 20, 30, 40, 50}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	/*
		 Iterar parecido al for each en otros lenguajes
		idenx: Posicion
		value: copia del elemento actual de la iteracion
	*/
	for index, value := range a {
		fmt.Printf("indice: %d, Valor: %d\n", index, value)
	}

	var b = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(b)

}

func rebanadas() {
	//Asi se define un slice, indice vacio.
	diasSemana := []string{"Domingo", "Lunes", "Martes",
		"Miercoles", "Jueves", "Viernes", "Sabado"}

	//Creamos otro slice a partir de un slice
	diaRebanada := diasSemana[0:5]

	fmt.Println(diaRebanada)
	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))

	//agregamos elementos al final del slice
	diaRebanada = append(diaRebanada, "Viernes", "Sabado", "Otro Dia")

	fmt.Println(diaRebanada)
	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))
	//envio dos rebanadas para que crear una nueva. borramos martes
	diaRebanada = append(diaRebanada[0:2], diaRebanada[3:]...)
	fmt.Println(diaRebanada)
	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))

}

func rebanadasMakeCopy() {
	//make permite crear slices mapas y canales
	nombres := make([]string, 5, 10)
	nombres[0] = "Cristian"
	fmt.Println(nombres)

	rebanada1 := []int{1, 2, 3, 4, 5}

	//Crea rebanada con una longitud inicial de
	rebanada2 := make([]int, 5)
	//copy metodo para copiar elementos de rebanad1 a rebanad 2, devuelve el número de elementos copiados.
	fmt.Println("Se copiaron:", copy(rebanada2, rebanada1))

	//Imprimir slices
	fmt.Println(rebanada1)
	fmt.Println(rebanada2)
}

func mapas() {
	colors := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}
	//Imprime el valor su existe
	fmt.Println(colors["rojo"])

	//Añade un nuevo clave, valor de color negro.
	colors["negro"] = "#000000"
	fmt.Println(colors)
	//verificar si existe en el mapa
	if valor, ok := colors["verde"]; ok {
		fmt.Println(valor)
	} else {
		fmt.Println("No existe la clave")
	}
	//Elimina azul de mapa
	delete(colors, "azul")
	fmt.Println(colors)

	// --- Iterar sobre un mapa usando 'for range' ---
	// El bucle 'for range' es la forma de recorrer todos los pares clave-valor de un mapa.
	for clave, valor := range colors {
		fmt.Printf("Clave: %s, Valor: %s\n", clave, valor)
	}
}
