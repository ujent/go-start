package main

import (
	"fmt"
)

func main() {
	names := [5]string{
		"Robert",
		"Jack",
		"Lize",
		"",
		"Arco",
	}

	fmt.Println(names)

	p := &names
	fmt.Println(p)

	a := names[0:4]
	b := names[3:]
	c := p[:]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	p[3] = "Nick"
	fmt.Println(p)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(names)
	fmt.Printf("\n")

	a[0] = "Daniel"
	fmt.Println(p)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(names)
	fmt.Printf("\n")

	var d []string
	fmt.Printf("cap = %d, len = %d\n", cap(d), len(d))

	if d == nil {
		fmt.Println("d == nil!")
	}

}
