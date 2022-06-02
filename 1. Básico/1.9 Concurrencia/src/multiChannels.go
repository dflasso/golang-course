package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	//crear un canal
	c := make(chan string, 2)

	//Ingreso de datos al canal
	c <- "Mensaje 1"
	c <- "Mensaje 2"

	/**
	len(c)  : cuantos datos o goroutines hay en channel
	cap(c)  : cantidad maxima de datos o goroutines puede tener el channel
	*/
	fmt.Println(len(c), cap(c))

	// Range y close
	/**
	solicita al runtime que cierre el canal

	Posterior a la ejecución no puede agregar más datos
	*/
	close(c)

	//error
	//c<-"Mensaje 3"

	// range permite recorrer arrays, slice, maps o canales
	for message := range c {
		fmt.Println(message)
	}

	/*
		Manejo multi canal
	*/
	// Select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensaje 1", email1)
	go message("mensaje 2", email2)

	// itera el número de canales que se quiere manejar
	for i := 0; i < 2; i++ {
		//
		select {
		//se almacena la salida del canal en m1
		case m1 := <-email1:
			fmt.Println("Email recibido de email 1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email 2", m2)
		}
	}
}
