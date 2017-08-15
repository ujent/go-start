package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	el1 := -1
	var el2 int
	var temp int

	return func() int {
		switch {
		case el1 == 0:
			{
				el2 = el1
				el1 = 1

				return el1
			}
		case el1 > 0:
			{
				temp = el1
				el1 = el2 + el1
				el2 = temp

				return el1
			}
		default:
			{
				el1 = 0
				return el1
			}
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
