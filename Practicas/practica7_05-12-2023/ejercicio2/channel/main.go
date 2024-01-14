package main

import (
	"fmt"
	"time"
)

type resultAndTime struct {
	label     string
	num       int
	timeTaked time.Duration
}

func main() {

	respch := make(chan resultAndTime)

	go func() {
		respch <- sumaPares()
	}()

	go func() {
		respch <- sumaImpares()

	}()

	resultado1 := <-respch
	resultado2 := <-respch
	fmt.Printf("%v suman %v el tiempo que le toma es %v\n", resultado1.label, resultado1.num, resultado1.timeTaked)
	fmt.Printf("%v suman %v el tiempo que le toma es %v\n", resultado2.label, resultado2.num, resultado2.timeTaked)

	close(respch)
}

/*go
# Ejercicio 2

Escriba un programa en donde haya dos gorutinas donde una suma los primeros 100 números pares y la otra los 100 primeros números impares. Se desea saber cual gorutina “gana”, es decir, quien termina termina primero.

Para sincronizar se puede usar grupos de espera o canales. Preferentemente podés realizar el ejercicio con ambos mecanismos para que te vayas familiarizando.
*/

func sumaPares() resultAndTime {
	startTime := time.Now()
	var pares int
	for i := 1; i < 20000; i++ {
		if i%2 == 0 {
			pares += i

		}
	}
	return resultAndTime{"Pares: ", pares, time.Since(startTime)}

}

func sumaImpares() resultAndTime {
	startTime := time.Now()
	var impares int
	for i := 1; i < 20000; i++ {
		if i%2 != 0 {
			impares += i
		}
	}
	return resultAndTime{"Impares: ", impares, time.Since(startTime)}

}
