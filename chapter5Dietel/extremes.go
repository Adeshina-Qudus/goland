package main

import "fmt"

func extremes() {
	sum := 0
	minimum := 0
	maximum := 0
	for counter := 0; counter < 5; counter++ {
		println("enter a number ")
		var input int
		fmt.Scan(&input)
		if counter == 0 {
			minimum = input
			maximum = input
		}
		if input > maximum {
			maximum = input
		}
		if input < minimum {
			minimum = input
		}
	}
	sum += minimum
	sum += maximum
	println("Minimum is : ", minimum)
	println("Maximum is : ", maximum)
	println("Sum Of Minimum And Maximum is  : ", sum)
}
