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
	"encoding/json"
	"fmt"
	"labora/clases/Practicas/practica12_22-12-2023/ejercicio1/numbers"
	"log"
	"net/http"
	"strconv"
)

type Num struct {
	Num int `json:"num"`
}

func main() {
	PORT := ":8080"
	log.Print("Running server on " + PORT)
	http.HandleFunc("/api/v1/sumEvens/", handlerEvens)
	http.HandleFunc("/api/v1/sumOdds/", handlerOdds)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func handlerEvens(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		num := r.URL.Query().Get("number")
		numInt, err := strconv.Atoi(num)
		if err != nil {
			w.Write([]byte("No se puede convertir el numero"))
			return
		}
		sum := numbers.SumaPares(numInt)
		sumStr := fmt.Sprint(sum)
		w.Write([]byte("La suma de los pares es: " + sumStr))

	}
	if r.Method == "POST" {
		var num Num
		err := json.NewDecoder(r.Body).Decode(&num)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.Write([]byte("No se puede convertir el numero"))
			return
		}
		var suma Num
		sum := numbers.SumaPares(num.Num)
		suma.Num = sum
		// sumStr := fmt.Sprint(sum)
		// w.Write([]byte("La suma de los pares es: " + sumStr))
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(suma)
	}
}
func handlerOdds(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		num := r.URL.Query().Get("number")
		numInt, err := strconv.Atoi(num)
		if err != nil {
			w.Write([]byte("No se puede convertir el numero"))
			return
		}
		sum := numbers.SumaImpares(numInt)
		sumStr := fmt.Sprint(sum)
		w.Write([]byte("La suma de los Impares es: " + sumStr))
	}
	if r.Method == "POST" {
		var num Num
		err := json.NewDecoder(r.Body).Decode(&num)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			w.Write([]byte("No se puede convertir el numero"))
			return
		}
		var suma Num
		sum := numbers.SumaImpares(num.Num)
		suma.Num = sum
		// sumStr := fmt.Sprint(sum)
		// w.Write([]byte("La suma de los Impares es: " + sumStr))
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(suma)
	}
}
