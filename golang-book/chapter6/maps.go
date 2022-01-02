package main

import "fmt"

func main() {
	// map: key-value pairs data type
	// same with slice, use `make()` function to declare map
	// no need to specify the length
	var x = make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])

	// "key" data type can be int
	var y = make(map[int]int)
	y[1] = 10
	// above declaration looks like array declaration
	// but map size can be changed
	// and map also not sequential, by writing that doesn't mean we have the value for y[0]
	fmt.Println(y[1])

	// build in function for map
	// `delete()`
	delete(y, 1)
	fmt.Println(y[1])

	elements := make(map[string]string)
	elements["He"] = "Helium"
	// if the value for a key is not exists:
	// 1. it will return 0 (if the key is int)
	// 2. it will return empty string (if the key is string)
	// accessing an element of a map will return two things:
	// 1. the value itself
	// 2. boolean value whether it's exists or not
	name, ok := elements["Ne"]
	fmt.Println(name, ok)
	if name, ok := elements["Ne"]; ok {
		// only be executed if it's true (exists)
		fmt.Println(name, ok)
	}
}
