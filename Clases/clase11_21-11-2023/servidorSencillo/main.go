package main

// Importar el paquete net/http:
import (
	"fmt"
	"net/http"
)

// Definir un manejador (handler):
func handler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("x-custom-header", "😅")
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("<h1>¡Hola, este es un servidor web simple en Go!</h1>"))
}

const serverPort string = "3005"

func main() {
	http.HandleFunc("/", handler) // asignar un manejador a un recurso (/ es el recurso "raiz")

	fmt.Println("Servidor esperando peticiones en puerto", serverPort)
	err := http.ListenAndServe(":"+serverPort, nil)

	if err != nil {
		fmt.Println("Error al iniciar servidor:", err)
	}
}
