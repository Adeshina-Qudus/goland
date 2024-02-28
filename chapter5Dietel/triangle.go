package main

import (
	"fmt"
)

func triangle() {

	for firstTriangle := 1; firstTriangle <= 10; firstTriangle++ {
		for asterisk := 1; asterisk <= firstTriangle; asterisk++ {
			fmt.Print("*")
		}
		fmt.Print(" ")
		for firstSpace := firstTriangle; firstSpace <= 10; firstSpace++ {
			fmt.Print(" ")
		}
		fmt.Print(" ")
		for secondTriangle := firstTriangle; secondTriangle <= 10; secondTriangle++ {
			fmt.Print("*")
		}
		fmt.Print(" ")
		for secondSpace := 1; secondSpace <= firstTriangle; secondSpace++ {
			fmt.Print(" ")
		}
		fmt.Print(" ")
		for thirdSpace := 1; thirdSpace <= firstTriangle; thirdSpace++ {
			fmt.Print(" ")
		}
		fmt.Print(" ")
		for thirdTriangle := firstTriangle; thirdTriangle <= 10; thirdTriangle++ {
			fmt.Print("*")
		}
		fmt.Print(" ")
		for fourthSpace := firstTriangle; fourthSpace <= 10; fourthSpace++ {
			fmt.Print(" ")
		}
		fmt.Print(" ")
		for fourthTriangle := 1; fourthTriangle <= firstTriangle; fourthTriangle++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
