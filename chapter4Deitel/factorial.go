package main

import "fmt"

func factorial() {
	fmt.Print("Enter number to calculate factorial : ")
	var number int
	fmt.Scan(&number)
	num := number
	for count := number - 1; count > 0; count-- {
		number *= count
	}
	println("Factorial of ", num, "is ", number)
}
