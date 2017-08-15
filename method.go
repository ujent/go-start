package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

type MyFloat float64

func (p Point) abs() float32 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func abs(p Point) float32 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (f MyFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

//we can do this only with our local types (declared in this packege only)
/*func (f float64) abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}*/

func main() {
	fmt.Println(Point{3, 4}.abs())
	fmt.Println(MyFloat(-10.7885).abs())

}
