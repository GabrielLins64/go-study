package main

import (
	"fmt"
	"yt_tutorials/m/19_embeddings/embeddings"
)

func main() {
	// embeddings.go
	var h embeddings.Human
	h = embeddings.Employee{Name: "John"}

	h.Breathe()
	h.Walk()
	h.Speak()
	
	// embeddings2.go
	d := embeddings.Dog{Age: 5}
	p1 := embeddings.Pet1{Name: "Milo", A: d}

	fmt.Println(p1.Name)
	p1.A.Breathe()
	p1.A.Walk()

	p2 := embeddings.Pet2{Name: "Oscar", Animal: d}
	fmt.Println(p2.Name)
	fmt.Println(p2.Animal)
	p2.Breathe()
	p2.Walk()
}