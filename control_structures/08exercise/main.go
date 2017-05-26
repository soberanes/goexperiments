package main

import (
	"fmt"
	"math"
)

// Round : round a number
func Round(f float64) float64 {
	return math.Floor(f + .5)
}

// RoundPlus : get the rounded number with digits given
func RoundPlus(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return Round(f*shift) / shift
}

// Sqrt : Get Newton's method square root
func Sqrt(x float64) float64 {
	var quotient float64
	var qtmp float64
	z := float64(1)

	for i := 0; i < 100; i++ {
		quotient = x / z
		z = (quotient + z) / x

		fmt.Println(RoundPlus(quotient, 6))
		fmt.Println(RoundPlus(z, 6))

		if quotient == qtmp {
			return z
		}

		qtmp = quotient
	}

	return 0
}

func main() {
	x := float64(2)
	sroot := Sqrt(x)
	if sroot == 0 {
		fmt.Println("Square root not found")
	} else {
		fmt.Printf("math sqrt (%f): %f\n", x, math.Sqrt(x))
		fmt.Printf("newton sqrt (%f): %f\n", x, sroot)
	}

}
