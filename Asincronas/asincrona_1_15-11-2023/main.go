package main

import (
	"fmt"
	"time"
)

func main() {
	var count1 *int
	count2 := new(int)
	countTemp := 5
	count3 := &countTemp
	t := &time.Time{}

	//Como obtener punteros
	fmt.Printf("count1: %#v\n", count1)
	fmt.Printf("count2: %#v\n", count2)
	fmt.Printf("count3: %#v\n", count3)
	fmt.Printf("time  : %#v\n", t)
	fmt.Println("----------------------------")
	//Como obtener el valor de los punteros
	//Para count 1, 2, y 3, tenemos que chequear que sea diferenente de nil y agregar * enfrente del nombre de la variable
	if count1 != nil {
		fmt.Printf("count1: %#v\n", *count1)
	}
	if count2 != nil {
		fmt.Printf("count2: %#v\n", *count2)
	}
	if count3 != nil {
		fmt.Printf("count3: %#v\n", *count3)
	}
	if t != nil {
		fmt.Printf("time  : %#v\n", *t)
		fmt.Printf("time  : %#v\n", t.String())
	}

	fmt.Println("----------------------------")

	//Prueba de las dos funciones
	var count int

	add5Value(count)
	fmt.Println("add5Value post:", count)

	add5Point(&count)
	fmt.Println("add5Point post:", count)

}

// Funciones para demostrar que sucede con el valor de count cuando es ingresado en un funcion como valor y como puntero
// 1. VALOR
func add5Value(count int) {
	count += 5
	fmt.Println("add5Value     :", count)
}

// 2. PUNTERO
func add5Point(count *int) {
	*count += 5
	fmt.Println("add5Point     :", *count)
}
