package main

import (
	"fmt"
	"time"
)

func salute() {
	now := time.Now().Hour()
	fmt.Printf("now: %g\n", now)
	switch {
	case now < 12:
		fmt.Println("Good morning!")
	case now < 17:
		fmt.Println("Good afternoon!")
	case now >= 17:
		fmt.Println("Good evening")
	default:
		fmt.Println("Good night!")
	}
}
func main() {
	today := time.Now().Weekday()
	fmt.Println(today)
	fmt.Println("When's Tuesday?")

	switch time.Tuesday {
	case today:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	default:
		fmt.Println("Too...")
	}

	fmt.Println("When's Saturday?")

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	salute()
}
