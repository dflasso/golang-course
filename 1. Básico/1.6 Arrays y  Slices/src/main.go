package main

import "fmt"

func main() {
	// Se le indica la cantidad máxima de elementos
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	// Slice. No se declara la cantidad máxima de elementos
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 11)
	fmt.Println(slice)

	// Append
	newSlice := []int{12, 13, 14}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	//recorer
	for index, value := range slice {
		fmt.Println(index, value)
	}
}
