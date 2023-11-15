package main

import "fmt"

func main() {
	fmt.Println(ejercicio6(1030))
	fmt.Println(ejercicio6(12045))
	fmt.Println(ejercicio6(176520))
}

/*
6. Para valientes: Pasar un periodo expresado en segundos a un periodo expresado en días, horas, minutos y segundos.
    1. 1030 segundos son 17 minutos con 10 segundos
    2. 12045 segundos son 3 horas, 20 minutos con 45 segundos
    3. 176520 segundos son 2 días, 1 hora, 2 minutos con 0 segundos.
*/

func ejercicio6(seg int) (int, int, int, int) {
	const segundosEnUnMinuto = 60
	const segundosEnUnaHora = 3600
	const segundosEnUnDia = 86400

	segundosRestantesDias := seg % segundosEnUnDia
	dias := (seg - segundosRestantesDias) / segundosEnUnDia

	segundosRestantesHoras := segundosRestantesDias % segundosEnUnaHora
	horas := (segundosRestantesDias - segundosRestantesHoras) / segundosEnUnaHora

	segundosRestantesMinutos := segundosRestantesHoras % segundosEnUnMinuto
	minutos := (segundosRestantesHoras - segundosRestantesMinutos) / segundosEnUnMinuto

	segundos := segundosRestantesMinutos

	return dias, horas, minutos, segundos
}
