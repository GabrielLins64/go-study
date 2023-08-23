package main

import (
	"fmt"
)

func main() {
	// Range em Slices
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// Range em Maps
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// Range em Strings
	for i, c := range "golang" {
		fmt.Printf("%d %c %d\n", i, c, c)
		// fmt.Println(i, c)
	}
}
