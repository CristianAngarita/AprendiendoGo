package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Función Connect que retorna un puntero a una conexión DB y un error (en caso de fallo)
func Connect() (*sql.DB, error) {
	// Cargar variables de entorno desde un archivo .env
	error := godotenv.Load()
	if error != nil {
		return nil, error
	}
	//Construcción del string de conexión (DSN: Data Source Name) desde el .env
	//"root:@tcp(localhost:3306)/db_contacts"
	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	// Abrir la conexión a la base de datos usando el driver MySQL
	db, err := sql.Open("mysql", dns)

	if err != nil {
		return nil, err
	}

	// Verificar que la conexión esté activa enviando con un ping
	if err := db.Ping(); err != nil {
		return nil, err
	}

	//Si todo funciona se envia esta respuesta.
	log.Println("Conexion exitosa\n")

	return db, nil
}
