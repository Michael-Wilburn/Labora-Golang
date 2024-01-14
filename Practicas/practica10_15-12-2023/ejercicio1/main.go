package main

import (
	"fmt"
	"sync"
	"time"
)

type AccesoSeguro struct {
	mu sync.Mutex
	sm map[int]int
}

func (a *AccesoSeguro) getOrCreate(key int, value int) int {
	a.mu.Lock()
	_, exist := a.sm[key]
	if !exist {
		a.sm[key] = value
	}
	defer a.mu.Unlock()
	return a.sm[key]
}
func main() {
	mapita := AccesoSeguro{sm: make(map[int]int)}
	go mapita.getOrCreate(1, 1)
	go mapita.getOrCreate(1, 2)
	go mapita.getOrCreate(1, 3)
	go mapita.getOrCreate(1, 4)

	time.Sleep(time.Millisecond * 10)

	mapita.mu.Lock()
	fmt.Printf("sm:%+v\n", mapita.sm)
	mapita.mu.Unlock()
}

/*
# Ejercicio 1
Entiendo el problema que se da cuando se quiere acceder de forma concurrente para modificar un mapa (tal como lo describe [este ejemplo](https://go.dev/play/p/ZT6Yb_6KGes)). Se desea implementar un tipo de datos abstracto (puede ser una struct) que represente un mapa sincronizado, es decir, que tenga accesos concurrentes seguros o que permite operaciones concurrentes en forma segura (thread safe).

1. ¿Es importante que sea concurrentemente seguro a nivel solo lectura?
2. ¿Qué diferencia observas si usas métodos mutadores que sean pointer receiver y value receiver? Los métodos mutadores son los que realizan cambios en el estado.
*/
