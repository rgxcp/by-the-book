package main

import "fmt"

var globalVariable string = "Hello World" // can't use shorthand syntax

func main() {
	variableOnlyWithinMain := "Hello World" // only accessible within nearest curly braces
	fmt.Println(variableOnlyWithinMain)
}

func another() {
	fmt.Println(globalVariable)
}
