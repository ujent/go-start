package main

import (
	"fmt"
)

func main() {
	a := make([]int, 4, 5)
	printSlice(a)

	b := append(a, 4)
	printSlice(a)
	printSlice(b)

	b = append(b, 2, 3, 5, 6, 6, 6, 9)
	printSlice(b)

	var c []int
	b = append(c, 3, 4)
	printSlice(c)
	printSlice(b)

	arr := [...]int{1, 3, 6, 10}
	fmt.Printf("cap = %d, len = %d, arr = %v\n", cap(arr), len(arr), arr)

	d := make([]int, len(b), (cap(b) * 2))
	count := copy(d, b)
	printSlice(d)
	fmt.Printf("Copied %d elements\n", count)

	//append slice to slice
	d = append(d, b...)
	printSlice(d)
	d = append(d, 1, 8, 7)
	printSlice(d)

	d = append(d, 9)
	printSlice(d)

	d = append(d, 10, 8, 7)
	printSlice(d)

	d = append(d, 1, 8, 7, 3, 4, 6)
	printSlice(d)

	d = make([]int, 2, 3)
	printSlice(d)
	d = append(d, 1, 2, 3, 4)
	printSlice(d)

	d = append(d, 1, 2, 3, 4)
	printSlice(d)

	var e []int
	e = append(e, d...)
	printSlice(e)
}

func printSlice(slice []int) {
	fmt.Printf("cap = %d, len = %d, %v\n", cap(slice), len(slice), slice)
}
