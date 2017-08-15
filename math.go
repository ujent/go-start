package main

import (
	"fmt"
	"math"
)

func pow(a, b, lim float64) float64 {
	if res := math.Pow(a, b); res < lim {
		return res
	} else {
		fmt.Printf("%g >= %g\n", res, lim)
		return lim
	}

}

func main() {
	fmt.Println(
		pow(10, 2, 100000),
		pow(4, 9, 77777777777),
		pow(34, 3, 233),
	)
}
