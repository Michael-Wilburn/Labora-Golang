/*
# Ejercicio 2

Escribir un programa que levante un servidor en el puerto 9090. El programa debe mantener en memoria una variable entera definida a nivel “global”. Además configurar el servidor para tener dos endpoints.

1. POST /api/v1/inc: Incrementa la variable entera en uno
2. POST /api/v1/dec: Decrementa la variable entera en uno
3. GET /api/v1/curr: Devuelve el valor de la variable entera.

Una vez escrito este programa servidor, vas a escribir un programa que juegue como cliente y tenga dos gorutinas. En una vas a enviar 10000 peticiones para incrementar y otra 1000 peticiones para decrementar. Finalmente vas a mandar una petición para obtener el valor de la variable entera una vez que terminen las dos gorutinas.

Se pide ejecutar el programa servidor y luego :

1. Ejecutando una vez el programa cliente
2. Ejecuta varias instancias del mismo programa cliente, es decir, ejecútalo varias veces en simultáneo.

Por último, se podrá apreciar que en el segundo caso para algún cliente el resultado que se le devuelve no es cero. ¿Esto a qué se debe? ¿ que hay que modificar para que a ambos clientes la petición “final” de obtener el valor de la variable entera devuelva un cero?
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var (
	globalInt    int
	globalIntMux sync.Mutex // Mutex para sincronizar el acceso a globalInt
)

const (
	incrementPath = "/api/v1/inc"
	decrementPath = "/api/v1/dec"
	currentPath   = "/api/v1/curr"
)

func main() {
	globalInt = 0 // Inicializa la variable global

	http.HandleFunc(incrementPath, incrementHandler)
	http.HandleFunc(decrementPath, decrementHandler)
	http.HandleFunc(currentPath, currentValueHandler)

	fmt.Println("Servidor escuchando en el puerto 9090...")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		panic(err)
	}
}

func incrementHandler(w http.ResponseWriter, r *http.Request) {
	globalIntMux.Lock()
	defer globalIntMux.Unlock()

	globalInt++
	fmt.Fprintf(w, "Variable incrementada a %d", globalInt)
}

func decrementHandler(w http.ResponseWriter, r *http.Request) {
	globalIntMux.Lock()
	defer globalIntMux.Unlock()

	globalInt--
	fmt.Fprintf(w, "Variable decrementada a %d", globalInt)
}

func currentValueHandler(w http.ResponseWriter, r *http.Request) {
	globalIntMux.Lock()
	defer globalIntMux.Unlock()

	// Crear una estructura para el JSON
	type Response struct {
		Value int `json:"value"`
	}

	// Crear la respuesta JSON
	response := Response{Value: globalInt}

	// Serializar la estructura Response a JSON
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Establecer el tipo de contenido de la respuesta como JSON
	w.Header().Set("Content-Type", "application/json")

	// Escribir el JSON en el cuerpo de la respuesta
	w.Write(jsonData)
}

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"sync/atomic"
// )

// var globalInt int32 // Usamos int32 para operaciones atómicas

// const (
// 	incrementPath = "/api/v1/inc"
// 	decrementPath = "/api/v1/dec"
// 	currentPath   = "/api/v1/curr"
// )

// func main() {
// 	globalInt = 0

// 	http.HandleFunc(incrementPath, incrementHandler)
// 	http.HandleFunc(decrementPath, decrementHandler)
// 	http.HandleFunc(currentPath, currentValueHandler)

// 	fmt.Println("Servidor escuchando en el puerto 9090...")
// 	if err := http.ListenAndServe(":9090", nil); err != nil {
// 		panic(err)
// 	}
// }

// func incrementHandler(w http.ResponseWriter, r *http.Request) {
// 	atomic.AddInt32(&globalInt, 1)
// 	fmt.Fprintf(w, "Variable incrementada a %d", atomic.LoadInt32(&globalInt))
// }

// func decrementHandler(w http.ResponseWriter, r *http.Request) {
// 	atomic.AddInt32(&globalInt, -1)
// 	fmt.Fprintf(w, "Variable decrementada a %d", atomic.LoadInt32(&globalInt))
// }

// func currentValueHandler(w http.ResponseWriter, r *http.Request) {
// 	val := atomic.LoadInt32(&globalInt)
// 	fmt.Fprintf(w, "Valor actual de la variable: %d", val)
// }

// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// const (
// 	incrementPath = "/api/v1/inc"
// 	decrementPath = "/api/v1/dec"
// 	currentPath   = "/api/v1/curr"
// )

// func main() {
// 	var incrementCh = make(chan int)
// 	var decrementCh = make(chan int)
// 	var currentCh = make(chan chan int)

// 	go func() {
// 		var globalInt int

// 		for {
// 			select {
// 			case inc := <-incrementCh:
// 				globalInt += inc
// 			case dec := <-decrementCh:
// 				globalInt -= dec
// 			case respCh := <-currentCh:
// 				respCh <- globalInt
// 			}
// 		}
// 	}()

// 	http.HandleFunc(incrementPath, func(w http.ResponseWriter, r *http.Request) {
// 		incrementCh <- 1
// 		fmt.Fprintf(w, "Variable incrementada")
// 	})

// 	http.HandleFunc(decrementPath, func(w http.ResponseWriter, r *http.Request) {
// 		decrementCh <- 1
// 		fmt.Fprintf(w, "Variable decrementada")
// 	})

// 	http.HandleFunc(currentPath, func(w http.ResponseWriter, r *http.Request) {
// 		respCh := make(chan int)
// 		currentCh <- respCh
// 		val := <-respCh
// 		fmt.Fprintf(w, "Valor actual de la variable: %d", val)
// 	})

// 	fmt.Println("Servidor escuchando en el puerto 9090...")
// 	if err := http.ListenAndServe(":9090", nil); err != nil {
// 		panic(err)
// 	}
// }
