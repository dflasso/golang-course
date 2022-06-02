package main

import (
	"fmt"
	"log"
	"strconv"
)

func isValidUser(userName, pass string) bool {
	return userName == "Alpha" && pass == "MyPassword"
}

func main() {

	// for condicional
	for i := 0; i < 10; i++ {
		//TODO
	}

	// for while
	counter := 0
	for counter < 10 {
		counter++
	}

	// for forever
	for {
		// Se ejecuta de por vida. Hasta que el programa termine su ejecución
		// debe existir una manera de terminar la ejecución
	}

	/*Operadores de comparación
	valor1 == valor2: Retorna TRUE si valor1 y valor2 son exactamente iguales.
	valor1 != valor2: Retorna TRUE si valor1 es diferente de valor2.
	valor1 < valor2: Retorna TRUE si valor1 es menor que valor2
	valor1 > valor2: Retorna TRUE si valor1 es mayor que valor2
	valor1 >= valor2: Retorna TRUE si valor1 es igual o mayor que valor2
	valor1 <= valor2: Retorna TRUE si valor1 es menor o igual que valor2.
	*/

	/*Operadores lógicos
	Operador AND:
		Este operador indica que todas las condiciones declaradas deben cumplirse para poderse marcar como TRUE. En Go, se utiliza este símbolo &&.

		Ejemplo1: 1>0 && 2 >0 Esto retornará TRUE porque tanto la primera como la segunda condición son verdaderas.

		Ejemplo2: 2<0 && 1 > 0 Esto retornará FALSE porque una de las condiciones no es verdadera.

	Operador OR:
		Este operador indica que al menos una de las condiciones debe cumplirse para marcarse como TRUE. En Go, se representa con el símbolo ||.

		Ejemplo: 2<0 || 1 > 0 Esto retornará TRUE porque la segunda condición se cumple, a pesar que la primera no.

	Operador NOT:
		Este operador retornará el opuesto al boleano que está dentro de la variable. Ejemplo:
		myBool :=  true
		fmt.Println(!myBool) // Esto retornará false
	*/

	num, err := strconv.Atoi("53")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(num)

	if num%2 == 0 {
		fmt.Println("Number is pair")
	} else {
		fmt.Println("Number is odd")
	}

	if isValidUser("Alpha5", "MyPassword") {
		fmt.Println("Credentials are valid")
	} else {
		fmt.Println("Credentials aren't valid")
	}

	// Estructura condicional - Switch
	var dia int8 = 4

	switch dia {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sabado")
	case 7:
		fmt.Println("Domingo")
	default:
		fmt.Println("Ese no es un día valido de la semana!")
	}

	// Estructura condicional - Switch
	var mes int8 = 10
	switch {
	case mes <= 3:
		fmt.Println("Primer Trimestre")
	case mes > 3 && mes <= 6:
		fmt.Println("Segundo Trimestre")
	case mes > 6 && mes <= 9:
		fmt.Println("Tercer Trimestre")
	case mes > 9 && mes <= 12:
		fmt.Println("Cuarto Trimestre")
	default:
		fmt.Println("Este no es un mes valido")
	}

}
