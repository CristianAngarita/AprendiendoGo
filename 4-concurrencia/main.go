package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	start := time.Now() //Hora exacta de ejecucion

	apis := []string{ //Lista de direecciones
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://apisomewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	ch := make(chan string) //Creamos un canal que solo permite strings

	//recorremos las apis
	for _, api := range apis {
		go checkApi(api, ch)
	}

	for i := 0; i < len(apis); i++ { //Este bucle es para "escuchar" las respuestas de todas nuestras goroutines.
		fmt.Print(<-ch)
	}

	elapsed := time.Since(start) // Calculamos cuánto tiempo ha pasado desde que empezamos.
	fmt.Printf("Listo, tomó %v segundos\n", elapsed.Seconds())
}

func checkApi(api string, ch chan string) { //verficamos cada api
	if _, err := http.Get(api); err != nil { // Intentamos hacer una solicitud HTTP a la 'api
		ch <- fmt.Sprintf("Error: %s esta caido\n", api) //en caso de error se envia la respuesta t se finaliza el goruntime
		return
	}
	ch <- fmt.Sprintf("SUCCESS: %s esta funcionnado\n", api) // Si no hay error se envia Succes
}
