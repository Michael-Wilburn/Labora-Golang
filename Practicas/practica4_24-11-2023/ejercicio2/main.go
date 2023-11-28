package main

import "fmt"

func main() {
	var arrayTest = []int{10, 20, 30, 40, 50}
	fmt.Println(borrar(arrayTest, 2))
	// fmt.Println(borrarElemento(arrayTest, 2))

}

/*
2. Haga que reciba un arreglo y una posición y “borre” el valor que hay en dicha posición. Por “borrar” entiendase que no se quita el elemento sino se mueven todos los elementos que están a la derecha un pasito a la izquierda, rellenado con el valor por defecto el extremo derecho. O sea…
    1. Si tengo el arreglo int[5]{10,20,30,40,50} y “borro” el elemento de la posición 2 (recordar que se indexan desde 0) quedaría ⇒ int[5]={10,20,40,50,0}
*/

// func borrar(array []int, index int) []int {

// 	for i, _ := range array {
// 		if i >= index && i < len(array)-1 {
// 			array[i] = array[i+1]
// 		}
// 	}
// 	array[len(array)-1] = 0
// 	return array

// }

func borrar(array []int, index int) []int {
	var nuevoArray []int
	for i, value := range array {
		if i == index {
			continue
		}
		nuevoArray = append(nuevoArray, value)
	}
	nuevoArray = append(nuevoArray, 0) // Añadir 0 al final después del bucle
	return nuevoArray
}

//Solucion alternativa
/*
func borrarElemento(arr []int, posicion int) []int {
	copy(arr[posicion:], arr[posicion+1:])
	arr[len(arr)-1] = 0
	return arr
}
*/
