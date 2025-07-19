package main

import (
	"bufio"
	"fmt"
	"go-mysql/database"
	"go-mysql/handlers"
	"go-mysql/models"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	/*
		//CREAR NUEVO
		newContact := models.Contact{
			Name:  "Nuevo usuario",
			Email: "nuevo@example.com",
			Phone: "123456789",
		}
		handlers.CreateContact(db, newContact) */

	/*
		//EDITAR CONTACTO
		newContact := models.Contact{
			Id:    6,
			Name:  "Nuevo usuario editado",
			Email: "nuevo@example.com",
			Phone: "123456789",
		}
		handlers.UpdateContact(db, newContact) */

	/* //LISTAR TODOS LOS CONTACTOS
	handlers.ListContact(db)
	//OBTENER CONTACTOS POR ID
	handlers.GetContactById(db, 6)
	//ELIMINAR CONTACTO
	handlers.DeleteContact(db, 6)
	//LISTAR TODOS LOS CONTACTOS
	handlers.ListContact(db)
	*/

	for {
		fmt.Println("\nMenú:")
		fmt.Println("1. Listar contactos")
		fmt.Println("2. Obtener contacto por ID")
		fmt.Println("3. Crear nuevo contacto")
		fmt.Println("4. Actualizar contacto")
		fmt.Println("5. Eliminar contacto")
		fmt.Println("6. Salir")
		fmt.Print("Seleccione una opción: ")

		//leer opcion
		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			handlers.ListContact(db)
		case 2:
			fmt.Print("Igrese Id del contacto")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.GetContactById(db, idContact)
		case 3:
			newContact := inputContactDetails(option)
			handlers.CreateContact(db, newContact)
		case 4:
			updateContact := inputContactDetails(option)
			handlers.UpdateContact(db, updateContact)
			handlers.ListContact(db)
		case 5:
			fmt.Print("Igrese Id del contacto que desea eliminar")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.DeleteContact(db, idContact)
		case 6:
			fmt.Println("Saliendo del programa .....")
			return
		default:
			fmt.Println("Opcion incorrecta.Ingrese opción valida")
		}
	}
}

// REcibe datos desde consolos y devuelve una instancia de models.Contact
// Se usa para crear nuevo o editar, depende de la opcion del usuario.
func inputContactDetails(option int) models.Contact {
	reader := bufio.NewReader(os.Stdin)

	var contact models.Contact

	//Si se quiere editar este debe hacer uso del id.
	if option == 4 {
		fmt.Print("Ingrese Id del contacto")
		var idContact int
		fmt.Scanln(&idContact)

		contact.Id = idContact
	}

	fmt.Print("Ingrese nombre del contacto: ")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)

	fmt.Print("Ingrese email del contacto: ")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)

	fmt.Print("Ingrese télefono del contacto: ")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact
}
