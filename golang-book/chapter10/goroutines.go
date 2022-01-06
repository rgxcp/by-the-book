package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() { // first goroutine
	// using the keyword `go` followed by a function call
	for i := 0; i < 10; i++ {
		go f(i) // second goroutine, run and return immediately to the next line without waiting for this function to finish
	}
	var input string // this line and below are needed, without it the program will exit immediately
	fmt.Scanln(&input)
}
