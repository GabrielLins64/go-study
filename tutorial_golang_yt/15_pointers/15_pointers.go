package main

import (
	"fmt"
)

func main() {
	i := 1
	fmt.Println("   [MAIN]: &i =", &i)
	fmt.Println("   [MAIN]: i =", i)
	
	zeroval(i)
	fmt.Println("   [MAIN]: i =", i)
	
	zeroptr(&i)
	fmt.Println("   [MAIN]: i =", i)
}

func zeroval(ival int) {
	ival = 0
	fmt.Println("[ZEROVAL]: ival =", ival)
	fmt.Println("[ZEROVAL]: &ival =", &ival)
}

func zeroptr(iptr *int) {
	*iptr = 0
	fmt.Println("[ZEROPTR]: iptr =", iptr)
	fmt.Println("[ZEROPTR]: &iptr =", &iptr)
	fmt.Println("[ZEROPTR]: *iptr =", *iptr)
}
