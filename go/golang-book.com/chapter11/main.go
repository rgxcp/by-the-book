package main

import (
	"fmt"

	m "../../golang-book/chapter11/math" // using "m" as an alias for "math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	fmt.Println(avg)
}
