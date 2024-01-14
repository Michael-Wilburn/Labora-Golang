/*
# Prácticas: Servers
# Ejercicio 1
Asumiendo que realizó los ejercicios para sumar los primeros n números pares y los primeros n números impares. Les voy a pedir que hagan lo siguiente:
1. Crear un proyecto golang donde las funciones sean parte de un paquete que no sea el main (Por ejemplo: puede llamarlo *numbers)* y que tenga un module path que sea un URL apuntando a un repositorio accesible.

2. Crearse otro proyecto golang para escribir un servidor que tenga dos endpoints
    1. /api/v1/sumEvens: Recibe como parámetro un número n y suma los primeros “n” números pares
    2. /api/v1/sumOdds: Recibe como parámetro un número n y suma los primeros “n” números impares
Se pide intencionalmente que NO se escriban las rutinas de código dentro del proyecto del servidor, sino que se importen usando los mecanismo de dependencias que tiene el lenguaje.
1. Se pide hacer ambos endpoints para recibir peticiones GET y POST. Si te parece más sencillo podés tener 4 endpoints, 2 para GET y 2 para POST.
*/

package main

import (
	"fmt"
	"labora/clases/Practicas/practica12_22-12-2023/ejercicio1/numbers"
	"log"
	"net/http"
	"strconv"
)

func main() {
	PORT := ":8080"
	log.Print("Running server on " + PORT)
	http.HandleFunc("/api/v1/sumEvens/", GetTaskFuncEvens)
	http.HandleFunc("/api/v1/sumOdds/", GetTaskFuncOdds)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

// Sumo los imapres
func GetTaskFuncOdds(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		n := r.URL.Path[len("/api/v1/sumOdds/"):]
		i, err := strconv.Atoi(n)
		if err != nil {
			w.Write([]byte("Esta tarea no se puede realizar"))
		} else {
			sum := numbers.SumaImpares(i)
			w.Write([]byte("La suma de los impares es: " + fmt.Sprint(sum)))
		}
	}
	if r.Method == "POST" {
		queryParameters := r.URL.Query()
		paramValue := queryParameters["number"]
		i, err := strconv.Atoi(paramValue[0])
		if err != nil {
			w.Write([]byte("Esta tarea no se puede realizar"))
		} else {
			sum := numbers.SumaImpares(i)
			w.Write([]byte("La suma de los impares es: " + fmt.Sprint(sum)))
		}
	}
}

// Sumo los pares
func GetTaskFuncEvens(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		n := r.URL.Path[len("/api/v1/sumEvens/"):]
		i, err := strconv.Atoi(n)
		if err != nil {
			w.Write([]byte("Esta tarea no se puede realizar"))
		} else {
			sum := numbers.SumaPares(i)
			w.Write([]byte("La suma de los pares es: " + fmt.Sprint(sum)))
		}
	}
	if r.Method == "POST" {
		queryParameters := r.URL.Query()
		paramValue := queryParameters["number"]
		i, err := strconv.Atoi(paramValue[0])
		if err != nil {
			w.Write([]byte("Esta tarea no se puede realizar"))
		} else {
			sum := numbers.SumaPares(i)
			w.Write([]byte("La suma de los pares es: " + fmt.Sprint(sum)))
		}
	}
}
