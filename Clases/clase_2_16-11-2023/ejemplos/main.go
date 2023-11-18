package main

import "fmt"

func main() {
	var n int
	fmt.Println("Ingrese un numero:")
	fmt.Scanf("%d", &n)

	// i := 1
	// for n > 0 { //loop o bucle
	// 	fmt.Println(i, "Hola")
	// 	n--
	// 	i++
	// }

	// for i := 1; n > 0; i++ {
	// 	fmt.Println(i, "hola")
	// 	n--
	// }

	// for {
	// 	if n == 0 {
	// 		break
	// 	}
	// 	n--
	// 	fmt.Println(i, "Hola")
	// 	i++
	// }

	sum := 0
	for i := 0; i < 37; i++ {
		sum += i
	}
	fmt.Println(sum)

}
