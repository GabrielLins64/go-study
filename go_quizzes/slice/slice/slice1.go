package slice

import "fmt"


func SliceOne() {
	fmt.Println("Slice 1")

	// We can infer the array length when using array literals with the ellipses ... operator.
	a := [...]int{0, 1, 2, 3}
	// Initially, the slice y is [2 3], and the slice x is [0] with three free element slots (in other words, its capacity is 4),
	x := a[:1]
	y := a[2:]
	// The 1st append call changes the slice x to [0 2 3] with one free element slot. And as the slice x and the array a share elements (before line 10), the array a is changed to [0 2 3 3]. The slice y and the array a also share elements, so the slice y is changed to [3 3].
	x = append(x, y...)
	// For the slice x has not enough free capacity, the 2nd append call allocates a new backing array for the slice x. So the call doesn't modify the array a. The slice x is changed to [0 2 3 3 3].
	x = append(x, y...)
	fmt.Println(a, x)
}
