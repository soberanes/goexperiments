package main

import "fmt"

/*
This function add y to x params. both are int and return an int.
function name(arg type [, ...]) return_type
*/
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
