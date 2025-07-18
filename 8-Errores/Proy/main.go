package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func saveContactasToFile(contact []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file) //Crea un codificador JSON que escribe en el archivo abierto.
	err = encoder.Encode(contact)    // Convierte el valor de contact en texto JSON y lo guarda directamente en el archivo.
	if err != nil {
		return err
	}
	return nil
}

func loadContacatsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file) //Crea un codificador JSON que escribe en el archivo abierto.
	err = decoder.Decode(contacts)   //Lee JSON desde archivo y lo convierte a estructura Go
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var contacts []Contact

	err := loadContacatsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error al cargar los contactos", err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("===Gestor de Contactos =====\n",
			"1. Agregar un contacto\n",
			"2. Mostrar todos los contactos\n",
			"3. Salir\n",
			"Elige una opcion:")

		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opcion", err)
			return
		}

		switch option {
		case 1:
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Telefono: ")
			c.Phone, _ = reader.ReadString('\n')

			contacts = append(contacts, c)

			if err := saveContactasToFile(contacts); err != nil {
				fmt.Println("Error al guardar el contacto", err)
			}
		case 2:
			fmt.Println("===========================================")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s Email %s Telefono: %s\n",
					index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("===========================================")
		case 3:
			return
		default:
			fmt.Println("Opción no valida.")
		}
	}

}
