package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5, 10)
	printSlice(a)

	a[4] = 6
	printSlice(a)

	a = a[1:]
	printSlice(a)

	a = a[:cap(a)]
	printSlice(a)

	a = a[:8]
	printSlice(a)
}

func printSlice(slice []int) {
	fmt.Printf("cap = %d, len = %d, %v\n", cap(slice), len(slice), slice)
}
