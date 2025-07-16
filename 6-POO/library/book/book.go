package book

import "fmt"

// Cualquier tipo que tenga un método PrintInfo() cumplirá esta interfaz.
type Printable interface {
	PrintInfo()
}

// Esto permite imprimir información sin importar el tipo concreto, si cumple con la interfaz.
func Print(p Printable) {
	p.PrintInfo()
}

type Book struct {
	title  string
	author string
	pages  int
}

//PrintInfo para Book. Es con receptor tipo puntero (*Book)
func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n", b.title, b.author, b.pages)
}

// Constructor para crear un nuevo Book.
// Devuelve un puntero al nuevo Book creado.
func NewBook(title string, author string, pages int) *Book {
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

// Método para modificar el título de un libro.
func (b *Book) SetTitle(title string) {
	b.title = title
}

// Método para obtener el título de un libro.
func (b *Book) GetTitle() string {
	return b.title
}

// Constructor para crear un nuevo TextBook.,Recibe parámetros para el Book y los nuevos campos.
// Se inicializa Book directamente dentro del struct
//Composicion con Book
type TextBook struct {
	Book
	editorial string
	level     string
}

func NewTextBook(title, author string, pages int, editorial, level string) *TextBook {
	return &TextBook{
		Book:      Book{title, author, pages},
		editorial: editorial,
		level:     level,
	}
}

//PrintInfo para Book. Es con receptor tipo puntero (*TextBook)
func (b *TextBook) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nEditorial: %s\nNivel: %s\n", b.title, b.author,
		b.pages, b.editorial, b.level)
}
