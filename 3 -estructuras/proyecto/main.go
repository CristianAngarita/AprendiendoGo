package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Una estructura es como un plantilla que sirve para:

1. modelar objetos del mundo real.
2. Organizar datos.
3. Reutilizar para crear tareas diferentes
4. Mejorar legibilidad del codigo
*/
type Tarea struct {
	nombre     string
	desc       string
	completado bool
}

/*
Definimos una struct que nos permite contener multiples Tareas.
Usa un slice de Tarea.
*/
type ListaTareas struct {
	tareas []Tarea
}

/*
Recibe una tarea individual
La función principal de este puntero (l *ListaTareas) es permitir
que el método agregarTarea modifique directamente la ListaTareas original
*/
func (l *ListaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

/*
	Metodo que marca una tarea como completado, recibe un indice para ello y lo

modifica a true como completada
*/
func (l *ListaTareas) marcarCompletado(index int) {
	l.tareas[index].completado = true
}

// Metodo para editar, recibe un indice y la tarea(nuevos datos) para ello
func (l *ListaTareas) editarTarea(index int, t Tarea) {
	l.tareas[index] = t
}

/*
Metodo para eliminar una tarea, utiliza un indice y se elimina por rebana, se toma la
posición por index.
*/
func (l *ListaTareas) eliminarTarea(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}

func main() {
	lista := ListaTareas{}
	leer := bufio.NewReader(os.Stdin)

	//Manejador de los metodos creados que permiten eliminar, editar, creal y actualizar tareas.
	for {
		var opcion int
		fmt.Println("Seleccione una opcion:\n",
			"1. Agregar tarea\n",
			"2. Marcar tarea completada\n",
			"3. Editar tarea.\n",
			"4. Eliminar tare.\n",
			"5 Salir")
		fmt.Print("ingrese la opción:")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			//Crea tarea nueva
			var t Tarea
			fmt.Println("Ingrese el nombre de la tarea:")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese la descripcion de la tarea:")
			t.desc, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Agregada correctamente")
		case 2:
			//Marca completada la tarea
			var index int
			fmt.Print("Ingrese el indice de la tarea que desea marcar como completada")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println("Marcada compleatada correctamente")
		case 3:
			//ACtualiza datos de una tarea
			var index int
			var t Tarea
			fmt.Print("Ingrese el indice de la tarea que desea actualizar")
			fmt.Scanln(&index)
			fmt.Println("Ingrese el nombre de la tarea:")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese la descripcion de la tarea:")
			t.desc, _ = leer.ReadString('\n')
			lista.editarTarea(index, t)
			fmt.Println("Actualizada correctamente")
		case 4:
			//Elimina una tarea
			var index int
			fmt.Print("Ingrese el indice para eliminar tarea")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
		case 5:
			//Termina ejecución del programa por consola
			fmt.Println("saliendo del progama")
			return
		default:
			fmt.Println("Opcion no valida")
		}

		//Lista las tareas y usa un for range para ello. Muestra indice y tarea
		fmt.Println("Lista de tareas:")
		fmt.Println("==================================================")
		for i, t := range lista.tareas {
			fmt.Printf("%d. %s - %s - Compleatado: %t\n", i, t.nombre, t.desc, t.completado)
		}
		fmt.Println("==================================================")
	}

}
