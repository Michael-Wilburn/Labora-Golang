/*
# Pair programming

Escribir un programa que declare una estructura llamada **`Persona`** que tenga los siguientes campos: **`nombre`** (cadena), **`edad`** (entero), **`altura`** (entero) y **`peso`** (entero).

A continuación, el programa deberá permitir el ingreso de los datos de 5 personas por teclado, en el orden que se desee.

Crear una función llamada **`ordenarPersonas`** que ordene las personas de acuerdo a un criterio de ordenamiento elegido por el usuario. El criterio de ordenamiento puede ser por nombre, edad, altura o peso, y el usuario debe seleccionarlo ingresando un número del 1 al 4.

La función deberá devolver el slice de personas ordenado según el criterio elegido.

Luego, crear otra función llamada **`buscarPersona`** que reciba el slice de personas y un valor a buscar, y devuelva la persona encontrada o **`nil`** si no se encuentra.

Deberás mostrar en pantalla: las personas creadas, sin ordenar, las personas ordenadas según el criterio elegido por el usuario, y el resultado de una búsqueda.

Además, se deberán aplicar las siguientes consideraciones:

- El programa debe verificar que los datos ingresados sean válidos y no permitir la creación de personas con valores incorrectos. Por ejemplo, no se permitirá ingresar una edad negativa o una altura igual a cero.
- La función de búsqueda debe permitir buscar por cualquier campo de la estructura Persona, no solo por nombre.
- Se deberá implementar una función que calcule el índice de masa corporal (IMC) de una persona, que se define como el peso de la persona dividido por el cuadrado de su altura (expresado en kg/m2).
- Al mostrar las personas, se deberá incluir su índice de masa corporal (IMC) y clasificarlo en las siguientes categorías:
  - Bajo peso: IMC menor a 18.5
  - Peso normal: IMC entre 18.5 y 24.9
  - Sobrepeso: IMC entre 25 y 29.9
  - Obesidad: IMC mayor a 30

- El programa deberá permitir al usuario realizar la opción de ordenar o buscar varias veces, sin tener que volver a ingresar los datos de las personas.
*/
package main

import (
	"fmt"
	"sort"
)

func main() {

	var personas = []Persona{{"tomas", 24, 1.80, 90}, {"pedro", 22, 1.50, 70}, {"raul", 21, 1.60, 75.5}}
	perona1 := personas[0]
	fmt.Println(ordenarPerson(personas, 0))
	fmt.Println(buscarPersona(personas, "22"))
	fmt.Println(perona1.calcuclarIMC())
	fmt.Println(perona1.clasificarIMC())

}

type Persona struct {
	nombre string
	edad   int
	altura float64
	peso   float64
}

func ordenarPerson(personas []Persona, crtierio int) []Persona {

	switch crtierio {
	case 1:
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].nombre < personas[j].nombre
		})
	case 2:
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].edad < personas[j].edad
		})
	case 3:
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].altura < personas[j].altura
		})
	case 4:
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].peso < personas[j].peso
		})
	default:
		return personas
	}

	return personas
}

func buscarPersona(personas []Persona, valor string) *Persona {

	for _, persona := range personas {
		if valor == persona.nombre || valor == fmt.Sprint(persona.edad) ||
			valor == fmt.Sprint(persona.altura) || valor == fmt.Sprint(persona.peso) {
			return &persona
		}
	}

	return nil
}

func (persona Persona) calcuclarIMC() float64 {
	imc := persona.peso / (persona.altura * persona.altura)

	return imc
}

func (persona Persona) clasificarIMC() string {
	imc := persona.calcuclarIMC()

	switch {
	case imc < 18.5:
		return "Bajo peso: IMC menor a 18.5"
	case imc >= 18.5 && imc < 25:
		return "Peso normal: IMC entre 18.5 y 24.9"
	case imc >= 25 && imc < 30:
		return "Sobrepeso: IMC entre 25 y 29.9"
	case imc > 30:
		return "Obesidad: IMC mayor a 30"
	default:
		return "IMC fuera de rango"
	}
}
