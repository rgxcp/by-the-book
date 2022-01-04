package main

import "fmt"

func zero(x int) {
	x = 0
}

// pointer refers to location in memory where the value is stored, rather than the value itself
// TL;DR => pointer points to something else
// pointer is represented using character `*`
// `xPtr *int` => xPtr is a pointer to an int
func zeroThatModified(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x is still 5 (it's not modified) since the x we passed is copied to the function

	// if we want to modify the original x, we can use special data type called pointer
	// `&` character used to find the address of a variable
	// it returns `*int` (pointer to an int)
	// `&x` in main and `xPtr` in `zeroThatModified` function refer to the same memory location
	zeroThatModified(&x)
	fmt.Println(x)

	// another way to get a pointer, using `new` function
	y := new(int)
	zeroThatModified(y)
	fmt.Println(*y)
}
