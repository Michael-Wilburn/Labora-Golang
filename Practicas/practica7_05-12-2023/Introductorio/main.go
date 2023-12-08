package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Suma secuencial
	startTime := time.Now()
	sumaSecuencial := sumaSecuencial(1, 100000000)
	timeTakedSecuancial := time.Since(startTime)

	fmt.Printf("Suma secuencial: %d\n", sumaSecuencial)
	fmt.Printf("Tiempo de ejecución secuencial: %s\n", timeTakedSecuancial)

	// Suma concurrente
	startTime = time.Now()
	sumaConcurrente := sumaConcurrente(1, 100000000)
	timeTakedConcurrency := time.Since(startTime)

	fmt.Printf("Suma concurrente: %d\n", sumaConcurrente)
	fmt.Printf("Tiempo de ejecución concurrente: %s\n", timeTakedConcurrency)

}

func sumaSecuencial(start, end int) int {
	var acc int
	for i := start; i <= end; i++ {
		acc += i
	}
	return acc
}

func sumaConcurrente(start, end int) int {
	var wg sync.WaitGroup
	result := 0
	respch := make(chan int)

	// Dividir el rango en dos partes y sumar cada parte concurrentemente
	middle := (start + end) / 2

	go func() {
		defer wg.Done()
		respch <- sumaSecuencial(start, middle)
	}()
	wg.Add(1)

	go func() {
		defer wg.Done()
		respch <- sumaSecuencial(middle+1, end)
	}()
	wg.Add(1)

	// Cerrar el canal después de que ambas goroutines hayan terminado
	go func() {
		wg.Wait()
		close(respch)
	}()

	// Recoger los resultados de las goroutines y sumarlos
	for value := range respch {
		result += value
	}

	return result

}

/*
	Ejercicio Introductorio

Hacer una suma de los primeros 100 numeros, de forma secuencial y luego de forma concurrente y apreciar la diferencia en el tiempo de ejecución.
*/
