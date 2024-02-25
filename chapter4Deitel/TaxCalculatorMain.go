package main

import "fmt"

func calculator() {
	tax := TaxCalculator{}

	for count := 1; count <= 3; count++ {
		fmt.Println("Enter name: ")
		var name string
		fmt.Scanln(&name)
		tax.setName(name)
		fmt.Println()

		fmt.Println("Enter earnings: ")
		var earnings float64
		fmt.Scanln(&earnings)
		tax.setEarnings(earnings)

		fmt.Println("Name of citizen is ", tax.getName())
		fmt.Println("Earnings of citizen is ", tax.getEarning())
		fmt.Println("Citizens total tax is ", tax.calculateTaxRate())
		fmt.Println()
	}
}
