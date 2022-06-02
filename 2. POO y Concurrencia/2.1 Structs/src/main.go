package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func NewEmployee(id int, name string) *Employee {
	return &Employee{
		id:   id,
		name: name,
	}
}

func main() {
	employee := Employee{}
	employee.id = 1
	employee.name = "Ana"
	employee.SetId(5)
	employee.SetName("Luis")

	/*Constructores*/
	e2 := Employee{
		id:   12,
		name: "Name 2",
	}

	e3 := new(Employee) // new devuelve un apuntador

	e4 := NewEmployee(1, "Name 4")

	fmt.Println(employee)
	fmt.Println(e2)
	fmt.Println(*e3) // con * referenciamos al valor.
	fmt.Println(*e4)

}
