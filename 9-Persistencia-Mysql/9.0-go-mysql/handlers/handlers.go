package handlers

import (
	"database/sql"
	"fmt"
	"go-mysql/models"
	"log"
)

// Handler definición : "un handler es una función que maneja una solicitud (request) y devuelve una respuesta (response)"

// Query se usa cuando se la consulta devuelve varias filas. Ejemplo SELECT
// QueryRow: Se usa cuando la consulta devuelve una sola fila.
// Exec: Cuando no se esperan filas, como en Delete, Update o Insert

// Función que listar todos los contactos
func ListContact(db *sql.DB) {
	query := "SELECT * FROM contact"

	//Ejectuar, devuelve una lista o un error.
	rows, error := db.Query(query)
	if error != nil {
		log.Fatal(error)
	}
	defer rows.Close()

	fmt.Println("LISTAS DE CONTACTOS")
	fmt.Println("-------------------------------------------------------------------------------------")

	// Itera sobre cada fila del resultado
	for rows.Next() {
		//Instancia de contact
		contact := models.Contact{}
		//Previene le error ya que esta columna Email permite los nulos.
		var valueEmail sql.NullString
		err := rows.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone) //Extrae los valores y los asigna por referencia
		if err != nil {
			log.Fatal(err)
		}
		// Si el email es válido y no está vacío, lo asigna. De lo contrario, asigna mensaje por defecto.
		if valueEmail.Valid && valueEmail.String != "" {
			contact.Email = valueEmail.String
		} else {
			contact.Email = "Sin correo electronico"
		}
		//Imprime los datos por consola.
		fmt.Printf("Id: %d, Nombre: %s, Email: %s, Telefono: %s\n",
			contact.Id, contact.Name, contact.Email, contact.Phone)
		fmt.Println("-------------------------------------------------------------------------------------")
	}
}

// Funcion listar contacto, recibe un id y un puntero a la conexion a la BD que imprimira UN contacto por ID
func GetContactById(db *sql.DB, contactID int) {
	query := "SELECT * FROM contact WHERE id = ?"

	row := db.QueryRow(query, contactID)

	contact := models.Contact{}
	var valueEmail sql.NullString

	//Extrae los valores y los asigna por referencia
	err := row.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)

	//Error en caso de que no se enceuntre el id asignado a un contacto en la base de datos
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("No se encontro ningun contacto con el ID %d", contactID)
		}
		log.Fatal(err)
	}
	// Si el email es válido y no está vacío, lo asigna. De lo contrario, asigna mensaje por defecto.
	if valueEmail.Valid && valueEmail.String != "" {
		contact.Email = valueEmail.String
	} else {
		contact.Email = "Sin correo electronico"
	}
	fmt.Println("CONTACTO")
	fmt.Println("-------------------------------------------------------------------------------------")
	fmt.Printf("Id: %d, Nombre: %s, Email: %s, Telefono: %s\n",
		contact.Id, contact.Name, contact.Email, contact.Phone)
	fmt.Println("-------------------------------------------------------------------------------------")
}

/*
Función para crear un nuevo contacto, recibe un puntero a la conexion de la
base de datos y una instancia del modelo Contact.
*/
func CreateContact(db *sql.DB, contact models.Contact) {

	query := "INSERT INTO contact (name, email, phone) VALUES (?, ?, ?)"

	//Ejecuta el query y recibe los 3 valores que se requieren, ?-> evita inyeccion SQL
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)

	//Manejo de error
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Contacto agregado exitosamente")
}

/*
Función para crear un nuevo contacto, recibe un puntero a la conexion de la
base de datos y una instancia del modelo Contact.
*/
func UpdateContact(db *sql.DB, contact models.Contact) {

	query := "UPDATE contact SET name = ?, email = ?, phone =? WHERE id = ?"

	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.Id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Contacto editado exitosamente")
}

/*
Función para eliminar un contacto, recibe un puntero a la conexion de la
base de datos y un ID
*/
func DeleteContact(db *sql.DB, contactID int) {
	query := "DELETE FROM contact WHERE id = ?"
	_, err := db.Exec(query, contactID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Contacto eliminado exitosamente")
}
