package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/health", health)

	fmt.Println("Servidor ejecutandose en htto://localhost:8080")

	http.ListenAndServe(":8080", nil)
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Está funcionando")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Hola desde la API en Go!")
}
