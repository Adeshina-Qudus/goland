package main

import "fmt"

func taxCalculator() {
	fmt.Print("Enter Name : ")
	var name string
	fmt.Scan(&name)
	fmt.Print("Enter Earnings : ")
	var earning float64
	fmt.Scan(&earning)
	var tax = 0.0
	if earning <= 30000 {
		tax = earning * 0.15
	} else {
		tax = earning * 0.20
	}
	fmt.Println(name, " Tax is ", tax)
}
