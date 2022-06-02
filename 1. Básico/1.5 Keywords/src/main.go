package main

import "fmt"

/*
Es recomendable para cerrar conexiones o archivos al finalizar un proceso
*/
func ejemploDefer() {
	//Defer
	/*Ejecuta la función antes de finalizar la función*/
	defer fmt.Println("hola")

	fmt.Println("mundo")
}

func main() {
	ejemploDefer()

	// Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		if i == 2 {
			fmt.Println("es 2")
			continue
		}

		if i == 8 {
			fmt.Println("termina el bucle")
			break
		}
	}
}
