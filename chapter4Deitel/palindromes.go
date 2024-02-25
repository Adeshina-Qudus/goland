package main

import "fmt"

func palindromes() {
	fmt.Print("Enter Number Consisting Of Five Digit : ")
	var inputtedDigit int
	fmt.Scan(&inputtedDigit)
	var number = inputtedDigit
	var reverse int
	var modulo int
	if inputtedDigit >= 10000 && inputtedDigit < 99999 {
		for inputtedDigit != 0 {
			modulo = inputtedDigit % 10
			reverse = reverse*10 + modulo
			inputtedDigit /= 10
		}
	} else {
		fmt.Print("Enter Number Consisting Of Five Digit : ")
		fmt.Scan(&inputtedDigit)
	}
	if number == reverse {
		fmt.Scan(number, "is a palindrome ")
	} else {
		fmt.Scan(number, " is not a palindrome ")
	}
}
