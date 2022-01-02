package main

import "fmt"

func main() {
	// same same but different with array
	// slice length is allowed to change
	var x []float64 // no need to define the length, by default the length is 0
	fmt.Println("x size:", len(x))

	// it's suggested to use `make()` function to create slice
	// slice with length of 5
	// third parameter is optional, it represents the capacity of the underlying array which the slice points to
	y := make([]float64, 5, 10)
	fmt.Println("y size:", len(y))

	// another way to create slice, using [low:high] expression
	arr := [5]float64{1, 2, 3, 4, 5}
	z := arr[1:4]
	fmt.Println("z:", z) // [2, 3, 4], not included the value of last index (high)
	// cheatsheets
	// arr[0:] == arr[0:len(arr)]
	// arr[:5] == arr[0:5]
	// arr[:]  == arr[0:len(arr)]

	// build in functions for slice
	// 1. `append()`
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	// 2. `copy()`
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)
}
