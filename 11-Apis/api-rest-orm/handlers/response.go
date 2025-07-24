package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func senData(rw http.ResponseWriter, data any, status int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)

	output, err := json.Marshal(data)
	if err != nil {
		http.Error(rw, "Error serializando JSON", http.StatusInternalServerError)
		return
	}

	_, _ = rw.Write(output)

	// Esto solo es útil para debugging
	fmt.Println(string(output))

}

func sendError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)

	fmt.Println(rw, "Resource Not found")
}
