package main

import "fmt"

func specificSum() {
	var total int
	var count int
	var firstInput int
	for {
		fmt.Print("Enter a number ")
		var input int
		fmt.Scan(&input)
		total += input
		if count == 0 {
			firstInput = input
			total = 0
		}
		count++
		if total == firstInput || total > firstInput {
			break
		}
	}

}
