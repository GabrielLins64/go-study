package embeddings

import "fmt"

type Dog struct {
	Age int
}

func (d Dog) Breathe() {
	fmt.Println("Dog breathes")
}

func (d Dog) Walk() {
	fmt.Println("Dog walk")
}

type Pet1 struct {
	A Animal
	Name string
}

type Pet2 struct {
	Animal
	Name string
}
