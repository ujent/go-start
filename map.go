package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func main() {
	m := make(map[int]int)
	m[1] = 2
	m[2] = 3
	fmt.Println(m)

	//var c map[string]int - cannot use without init, can't add values in nil
	var c map[string]int
	fmt.Println(c)

	var b = make(map[string]int)
	fmt.Println(b)
	b[""] = 0
	b["tt"] = 1
	fmt.Println(b)

	d := map[int]int{
		1: 8,
		2: 9,
		7: 4,
	}
	fmt.Println(d)

	//it's possible to omit type Point
	e := map[int]Point{
		1: {1, 2},
		2: {2, 3},
		3: Point{6, 9},
	}

	fmt.Println(e)
	e[1] = Point{6, 7}
	delete(e, 2)
	el, ok := e[2]
	fmt.Println(e)
	fmt.Printf("elem: %T%v, ok: %v\n", el, el, ok)
}
