package main

import (
	"log"
	"net/http"
	"rpsweb/handlers"
)

func main() {
	//crea enrutador para manejar rutas
	router := http.NewServeMux()

	//Manejador para servir los archivos estaticos
	fs := http.FileServer(http.Dir("static"))
	//Ruta para acceder a los archivos estaticos
	router.Handle("/static/", http.StripPrefix("/static/", fs))
	//configurar ruta
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	port := ":8082"
	log.Printf("SERVIDOR escuchando en http://localhost%s\n:", port)
	log.Fatal(http.ListenAndServe(port, router))
}
