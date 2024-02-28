package main

import "fmt"

func diamond() {
	for row := 0; row < 10; row++ {
		for column := row; column < 10; column++ {
			fmt.Print(" ")
		}
		for column := 1; column < row; column++ {
			fmt.Print("*")
		}
		for column := 0; column < row; column++ {
			fmt.Print("*")
		}

		fmt.Println("")
	}
	for row := 0; row < 10; row++ {
		for column := 0; column < row; column++ {
			fmt.Print(" ")
		}
		for column := row + 1; column < 10; column++ {
			fmt.Print("*")
		}
		for column := row; column < 10; column++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
