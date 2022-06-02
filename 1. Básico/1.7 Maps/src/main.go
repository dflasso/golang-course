package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["jose"] = 14
	m["pedro"] = 20

	//recorrer map
	for key, value := range m {
		fmt.Println(key, value)
	}

	//encontrar valor
	data, data_exist := m["jose"]
	fmt.Println(data, data_exist)
}
