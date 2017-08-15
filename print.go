//operations with print
package main

import (
	"fmt"
	"os"
	"strings"
)

func main0() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "; "
	}

	fmt.Println(len(os.Args))
	fmt.Println("args: " + s)
}

func main2() {
	s, sep := "", ""
	arr := [3]string{"a1", "a2", "a3"}
	for _, arg := range arr[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

func main() {
	fmt.Println(strings.Join(os.Args[1:], "; "))
}

func main3() {
	arr := [3]string{"a1", "a2", "a3"}
	fmt.Println(strings.Join(arr[0:], "; "))
}
