package main

import "fmt"

func main() {
	// similar to while loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1 // can be shortened to `i++`
	}

	// the classic...
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
