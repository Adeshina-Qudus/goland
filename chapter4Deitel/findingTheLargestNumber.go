package main

import "fmt"

func largestNumber() {
	var largest int
	for count := 1; count <= 10; count++ {
		fmt.Print("Enter a number ")
		var input int
		fmt.Scan(&input)
		if input > largest {
			largest = input
		}
	}
	fmt.Print("The largest is ", largest)
}
