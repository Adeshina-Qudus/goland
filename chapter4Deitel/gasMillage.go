package main

import "fmt"

func gasMillage() {
	var averageOfMiles float64
	var total int
	var count int
	fmt.Print("Enter miles driven : ")
	var milesValue int
	fmt.Scan(&milesValue)
	fmt.Print("Enter Gallon Used : ")
	var gallonValue int
	fmt.Scan(&gallonValue)
	for {
		total += milesValue / gallonValue
		count++
		fmt.Print("do you wish to continue : ")
		var question string
		fmt.Scan(&question)
		if question == "no" {
			break
		}
		fmt.Print("Enter miles driven : ")
		fmt.Scan(&milesValue)
		fmt.Print("Enter Gallon Used : ")
		fmt.Scan(&gallonValue)
	}
	averageOfMiles = float64(total / count)
	fmt.Print("The Combine Miles Per gallon is : ", total, "\n",
		"The Miles Per gallon is : ", averageOfMiles)
}
