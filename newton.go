package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := 1.0
	var index int

	for i := 1; i < 10; i++ {
		z = z - (math.Pow(z, 2)-x)/2*z
		index = i
	}

	fmt.Println(index)
	return z
}

func main() {
	newton := sqrt(2)
	common := math.Sqrt(2)
	fmt.Println(newton)
	fmt.Println(common)
	fmt.Printf("delta = %g", common-newton)
}
