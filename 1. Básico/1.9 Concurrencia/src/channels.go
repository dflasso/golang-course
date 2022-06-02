package main

import "fmt"

/**
@param c chan<- string  (canal solo de entrada de datos)
@param c <-chan string  (canal solo de sálida de datos)
*/
func say(text string, c chan<- string) {
	// Ingresamos el dato al canal
	c <- text
}

func main() {
	/*
		@param tipo de dato que va a pasar por el canal
		@param num de datos que van a pasar por el canal

		Es una buena practica especificar el número de datos que va a manejar el canal
		Si no se especifica el canal manejara 'n' cantidad de datos, de forma dinamica
	*/
	c := make(chan string, 1)

	fmt.Println("Hello")

	// inicia la goroutine
	go say("Bay", c)

	//Sacamos el dato del canal
	fmt.Println(<-c)
}
