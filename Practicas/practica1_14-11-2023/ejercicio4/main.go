package main

import (
	"fmt"
)

func main() {
	var dias int
	fmt.Scan(&dias)

	var horas int
	fmt.Scan(&horas)

	var minutos int
	fmt.Scan(&minutos)

	var segundos int
	fmt.Scan(&segundos)

	fmt.Println(ejercicio4(dias, horas, minutos, segundos))

}

//4.Redactar un algoritmo para pasar un periodo expresado en días, horas, minutos y segundos a periodo expresado en segundos.

func ejercicio4(dias int, horas int, minutos int, segundos int) int {
	const segundosEnUnMinuto = 60
	const segundosEnUnaHora = 3600
	const segundosEnUnDia = 86400

	totalSegDias := dias * segundosEnUnDia
	totalSegHoras := horas * segundosEnUnaHora
	totalSegminutos := minutos * segundosEnUnMinuto

	totalSegundos := totalSegDias + totalSegHoras + totalSegminutos + segundos

	return totalSegundos

}
