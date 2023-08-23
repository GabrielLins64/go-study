package main

import (
	"fmt"
)

func main() {
	nextInt := intSeq()
	fmt.Println("nextInt chamada:", nextInt())
	fmt.Println("nextInt chamada:", nextInt())
	fmt.Println("nextInt chamada:", nextInt())
	fmt.Println()
	
	newInts := intSeq()
	fmt.Println("newInts chamada:", newInts())
	fmt.Println("newInts chamada:", newInts())
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
