package main

import (
	"fmt"
)

func main() {
	fmt.Println("1 + 2 =", plus(1,2))
	fmt.Println("1 + 2 + 3 =", plusPlus(1,2,3))

	a, b := vals()
	fmt.Println(a, b)
	_, c := vals()
	fmt.Println(c)

	res := sum(2,3,4,5,6)
	fmt.Println("res:", res)

	nums := []int{1,2,3,4}
	res2 := sum(nums...) // Operador de destructuring
	fmt.Println("res2:", res2)
}

func plus(a int, b int) int {
	return a + b
}

// Aqui, todos os argumentos são considerados "int"
func plusPlus(a, b, c int) int {
	return a + b + c
}

// Retorno de múltiplos valores
func vals() (int, int) {
	return 7, 11
}

// Número arbitrário de argumentos
func sum(nums ...int) int {
	fmt.Print("args: ")
	fmt.Print(nums, " ")
	fmt.Println()
	var total int = 0
	for _, num := range nums {
		total += num
	}
	return total
}
