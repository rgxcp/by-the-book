package main

import "fmt"

func main() {
	var x string = "first"
	fmt.Println(x)

	x = "second"
	fmt.Println(x)

	var y = "third"
	fmt.Println(y)

	fourthVariable := "fourth" // preferable
	fmt.Println(fourthVariable)

	var (
		a = 5
		b = 10
		c = 15
	) // shorthand for creating multiple variables at once
	fmt.Println(a, b, c)
}
