// Go suporta a incorporação (Embedding) de estruturas e interfaces 
// para expressar uma composição de tipos mais perfeita. Incorporando 
// Interface em Outra Interface. Uma interface pode embutir qualquer 
// número de interfaces nela, assim como pode ser embutida em qualquer 
// interface. Todos os métodos da interface incorporada tornam-se 
// parte da interface.É uma maneira de criar uma nova interface 
// mesclando algumas pequenas interfaces.

package embeddings

import (
	"fmt"
)

type Animal interface {
	Breathe()
	Walk()
}

type Human interface {
	Animal
	Speak()
}

type Employee struct {
	Name string
}

func (e Employee) Breathe() {
	fmt.Println("Employee breathes")
}

func (e Employee) Walk() {
	fmt.Println("Employee walks")
}

func (e Employee) Speak() {
	fmt.Println("Employee speaks")
}
