package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

func newPerson(name string, age int) *person {
	p := person{name: name}
	p.age = age
	return &p
}

func main() {
	fmt.Println(person{"Ana", 23})
	fmt.Println(person{name: "Alice", age: 30})
	var newP *person = newPerson("Gabs", 27)
	fmt.Println(*newP)
	fmt.Println("name:", (*newP).name)
	fmt.Println("name:", newP.name) // same as above
	newP.name = "John"
	fmt.Println("name:", newP.name)
}
