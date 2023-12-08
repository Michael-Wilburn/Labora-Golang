package main

import (
	"fmt"
	"sync"
	"time"
)

type resultAndTime struct {
	label     string
	num       int
	timeTaked time.Duration
}

func main() {

	var wg sync.WaitGroup
	respch := make(chan resultAndTime)

	go func() {
		defer wg.Done()
		respch <- sumaPares()
	}()
	wg.Add(1)

	go func() {
		defer wg.Done()
		respch <- sumaImpares()
	}()
	wg.Add(1)

	go func() {
		wg.Wait()
		close(respch)
	}()

	for value := range respch {
		fmt.Printf("%v suman %v el tiempo que le toma es %v\n", value.label, value.num, value.timeTaked)

	}

}

/*go
# Ejercicio 2

Escriba un programa en donde haya dos gorutinas donde una suma los primeros 100 números pares y la otra los 100 primeros números impares. Se desea saber cual gorutina “gana”, es decir, quien termina termina primero.

Para sincronizar se puede usar grupos de espera o canales. Preferentemente podés realizar el ejercicio con ambos mecanismos para que te vayas familiarizando.
*/

func sumaPares() resultAndTime {
	startTime := time.Now()
	var pares int
	for i := 0; i < 200; i++ {
		if i%2 == 0 {
			pares += i

		}
	}
	return resultAndTime{"Pares: ", pares, time.Since(startTime)}

}

func sumaImpares() resultAndTime {
	startTime := time.Now()
	var impares int
	for i := 0; i < 200; i++ {
		if i%2 != 0 {
			impares += i
		}
	}
	return resultAndTime{"Impares: ", impares, time.Since(startTime)}

}
