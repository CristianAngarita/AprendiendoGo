package animal

import "fmt"

type Animal interface {
	Sonido()
}

// Esta interfaz requiere que cualquier tipo que la implemente tenga un método llamado Sonido().
type Perro struct {
	Nombre string
}

// El método Sonido() está definido con un receptor *Perro (puntero a Perro).
func (p *Perro) Sonido() {
	fmt.Println(p.Nombre + " ladra")
}

// Definimos un struct llamado Gato que también tiene un campo Nombr
type Gato struct {
	Nombre string
}

// El método Sonido() con receptor *Gato imprime el sonido característico de un gato
func (g *Gato) Sonido() {
	fmt.Println(g.Nombre + " hace miau")
}

//función que acepta cualquier tipo que implemente la interfaz Animal.
// Llama al método Sonido() sin importar si es un perro, gato u otro animal que implemente la interfaz.
func HacerSonido(animal Animal) {
	animal.Sonido()
}
