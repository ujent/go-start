package main

import (
	"fmt"
)

func main() {
	i := 0

	defer fmt.Println(i)

	for j := 0; j < 10; j++ {
		i = i + j + 5
		defer fmt.Println(i)
	}

	fmt.Println(i)
}
