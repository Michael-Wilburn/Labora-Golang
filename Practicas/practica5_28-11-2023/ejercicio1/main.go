package main

import "fmt"

func main() {
	var persona1 Persona = Persona{
		nombre:   "John Doe",
		edad:     31,
		ciudad:   "New York",
		telefono: 5554444,
	}

	var persona2 Persona = Persona{
		nombre:   "Jane Doe",
		edad:     30,
		ciudad:   "Gotica",
		telefono: 5556666,
	}

	fmt.Println(persona1)
	fmt.Println(persona2)

	persona2.setCambiarCiudad("Gotica")

	fmt.Println(persona1)
	fmt.Println(persona2)

	persona2.setCambiarCiudad("Miami")

	fmt.Println(persona1)
	fmt.Println(persona2)

}

type Persona struct {
	nombre   string
	edad     int
	ciudad   string
	telefono int
}

func (persona *Persona) setCambiarCiudad(ciudad string) string {
	if ciudad != persona.ciudad {
		persona.ciudad = ciudad
		return ciudad
	}
	return persona.ciudad
}
