package main

import (
	"fmt"
	"os"
)

// functions are built up in a "stack" (that one video on Insinyur Online)
func average(scores []float64) float64 {
	total := 0.0
	for _, value := range scores {
		total += value
	}
	return total / float64(len(scores))
}

// function with named return type, in this case `r`
func f1() (r int) {
	r = 1
	return // no need to specify the variable that need to be returned
}

// function that return multiple variables
// the second variable often used to hold errors
// res, err := something()
func f2() (int, int) {
	return 7, 5
}

// `...` keyword before the `int` data type means the parameter `args` can hold 0 or `n` values
// it's called variadic function
// and can only be used at the last parameter
func add(args ...int) int {
	total := 0
	for _, value := range args {
		total += value
	}
	return total
}

// closure
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// recursion
// function that is able to call itself
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
	// 5 * 4 * 3 * 2 * 1
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func main() {
	scores := []float64{98, 93, 77, 82, 83}
	average := average(scores)
	fmt.Println(average)

	fmt.Println(f1())

	x, y := f2()
	fmt.Println(x, y)

	fmt.Println(add(1, 2, 3))
	z := []int{1, 2, 3}
	fmt.Println(add(z...)) // can also pass a slice of int, with `...` suffix

	// function inside a function
	// `increment` is a variable that has `func` data type
	current := 0
	increment := func() int {
		// `increment` function inside `main` function and use `current` variable (non-local) form a closure
		current++
		return current
	}
	fmt.Println(increment())
	fmt.Println(increment())

	// another way to use closure
	// `makeEvenGenerator` is a function that return a function
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4

	// closure and recursion form the basis of what-so-called functional programming paradigm
	// it's harder than OOP
	// it's very mathematical
	// that's why engineers are called the oompa loompa of science
	fmt.Println(factorial(5))

	// defer
	// special statement schedules a function call to be run after the function completes
	defer second()
	first()
	// defer often use when resources need to be freed in some way, closing file for example
	// defer still run even if runtime panic occurs
	file, _ := os.Open("README.md")
	defer file.Close()

	// recover
	// handle runtime panic
	// it stop the panic and returns the value that was passed to the call to panic
	// need to be paired with defer, since if we use the normal way, panic will immediately stops execution
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	// panic generally indicates a programmer error
	// - accessing index of an array that's out of range
	// - forget to initialize a map, etc.
	panic("PANIC")
}
