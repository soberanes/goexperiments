package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var z = uint(f) //omit explit type declaration in var z uint
	fmt.Println(x, y, z)
}
