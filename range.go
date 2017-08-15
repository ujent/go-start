package main

import (
	"fmt"
)

func main() {
	sl := []int{1, 2, 3, 10, 22}

	for i, el := range sl {
		fmt.Printf("index: %d, element: %d\n", i, el)
	}

	// if index isn't necessary
	for _, el := range sl {
		fmt.Printf("el: %v\n", el)
	}
}
