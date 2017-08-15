package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		[]int{0, 1, 1},
		[]int{1, 1, 0},
		[]int{1, 0, 0},
	}

	for i := 0; i < len(matrix); i++ {
		fmt.Printf("%d\n", matrix[i])
	}
}
