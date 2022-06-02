package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgumento(a, b int, message string) {
	fmt.Println(a, b, message)
}

func returnOneValue(a, b int) int {
	return a * b
}

func returnManyValue(a, b int) (c, d int) {
	return a * b, a + b
}

func main() {
	normalFunction("Hola Mundo")
	tripleArgumento(1, 2, "hola mundo")

	result := returnOneValue(4, 5)
	fmt.Println(result)

	resultA, resultB := returnManyValue(6, 5)
	fmt.Println(resultA, resultB)

	// decartar un valor de retorno
	resultFirst, _ := returnManyValue(6, 5)
	fmt.Println(resultFirst)
}
