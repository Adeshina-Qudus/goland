package main

import (
	"fmt"
	"strconv"
)

func barchart() {
	var holdingInputtedValues string
	for count := 1; count <= 5; count++ {
		println("ENTER A NUMBER BETWEEN ONE AN THIRTY ")
		var userInput int
		fmt.Scan(&userInput)
		for userInput < 1 || userInput > 30 {
			count--
			println("ENTER A VALID NUMBER FROM ONE TO THIRTY")
			fmt.Scan(&userInput)
		}
		holdingInputtedValues += strconv.Itoa(userInput)
	}
	println()
	for _, value := range holdingInputtedValues {
		count, _ := strconv.Atoi(string(value))
		for counter := 0; counter < count; counter++ {
			print("*")
		}
		println()
	}

}
