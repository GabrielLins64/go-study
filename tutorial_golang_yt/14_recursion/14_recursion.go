package main

import (
	"fmt"
)

func main() {
	num := 6
	fmt.Printf("fact(%d) = %d\n", num, fact(num))

	// Declara uma variável do tipo função anônima (closure)
	// que recebe um parâmetro do tipo inteiro e retorna
	// um inteiro
	var fib func(n int) int
	// use a variável criada anteriormente para receber o
	// retorno da implementação da função anônima fib
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Printf("fib(%d) = %d\n", num, fib(num))
}

func fact(n int) int {
	if n < 2 {
		return 1
	}
	return n * fact(n-1)
}
