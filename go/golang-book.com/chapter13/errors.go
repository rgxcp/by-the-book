package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("Custom error message")
	fmt.Println(err)
}
