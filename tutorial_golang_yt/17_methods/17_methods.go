package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

func (r *rect) area() int {
	// fmt.Printf(" [area]: r -> %p\n", r)
	return r.width * r.height
}

// A função acima é mais eficiente em termos de memória devido
// a, na hora da chamada do método, o de cima usar um ponteiro
// para o rect, enquanto a função abaixo faz uma cópia (passa-
// gem por valor).
func (r rect) perim() int {
	// fmt.Printf("[perim]: &r -> %p\n", &r)
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println(" [main]: r.area() ->", r.area())
	fmt.Println(" [main]: r.perim() ->", r.perim())

	rp := &r
	fmt.Println(" [main]: rp.area() ->", rp.area())
	fmt.Println(" [main]: rp.perim() ->", rp.perim())
}
