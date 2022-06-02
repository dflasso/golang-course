package main

import (
	"fmt"
	"strconv"
)

func main() {
	// con error
	//myValue, err := strconv.ParseInt("NaN", 0, 64)

	// sin error
	myValue, err := strconv.ParseInt("7", 0, 64)

	// Validando errores.
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Println(myValue)
	}

	// Mapa clave valor.
	m := make(map[string]int)
	m["key"] = 6
	fmt.Println(m["key"])

	// Slice de enteros.
	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}
	s = append(s, 16)
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}

	//punteros
	g := 25
	fmt.Println(g) // imprime el valor entero 25
	h := &g
	fmt.Println(h) // imprimer la direccion de memoria.
	i := *h
	fmt.Println(i) // Imprime el valor por de g
}
