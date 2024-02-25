package main

import "fmt"

func checkerBoard() {
	for outterloop := 1; outterloop <= 8; outterloop++ {
		if outterloop%2 == 0 {
			fmt.Print(" ")
		}
		for innerloop := 1; innerloop <= 8; innerloop++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	fmt.Println()
}
