package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	var res = make(map[string]int)
	for _, word := range strings.Fields(s) {
		el, ok := res[word]

		if ok {
			res[word] = el + 1
		} else {
			res[word] = 1
		}

	}
	return res
}

func main() {
	s := "I ate a donut. Then I ate another donut."
	r := WordCount(s)

	fmt.Println(r)
}
