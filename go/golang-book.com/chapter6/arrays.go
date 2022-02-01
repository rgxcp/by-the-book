package main

import "fmt"

func main() {
	var x [5]float64 // fixed length
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	// shorthand syntax for creating array and fill its value directly
	// x := [5]float64{98, 93, 77, 82, 83}

	// it can also be break into multiple line
	// extra trailing `,` is required in multiple line declaration
	// x := [5]float64{
	// 	98,
	// 	93,
	// 	77,
	// 	82,
	// 	83,
	// }

	var total float64 = 0 // the type declaration is required, since by default it will typed to int
	// Ruby `.each` like loop
	// the first parameter can hold the current position in array
	// since it's not used, keyword `_` is used
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x))) // type convertion, from int to float64
}
