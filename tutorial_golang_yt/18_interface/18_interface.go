// Uma interface é uma estrutura/sintaxe de programação que permite ao programador impor certas propriedades em um objeto (classe). São coleções nomeadas de assinaturas de métodos.

package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// Métodos do rect
func (r *rect) area() float64 {
	return r.width * r.height
}

func (r *rect) perim() float64 {
	return r.width*2 + r.height*2
}

// Métodos do circle
func (c *circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c *circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Chamando métodos da interface
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Área:", g.area())
	fmt.Println("Perímetro:", g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	fmt.Print("Medidas do retângulo: ")
	measure(&r)
	fmt.Print("Medidas do círculo: ")
	measure(&c)
}
