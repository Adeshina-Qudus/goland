package main

import "fmt"

func SalesCommissionMain() {
	sales := SalesCommission{}
	var totalOfItemSold float64 = 0
	var condition bool = true
	for condition {
		fmt.Println("Enter price for item sold: ")
		var priceOfItems float64
		fmt.Scanln(&priceOfItems)
		totalOfItemSold += priceOfItems
		fmt.Println("Do you wish to continue: ")
		var answer string
		fmt.Scanln(&answer)
		if answer == "no" {
			condition = false
		}
	}
	sales.setPriceForItemSold(totalOfItemSold)
	fmt.Println("Total of item sold is ", sales.getPriceForItemSold())
	fmt.Println("Total earnings is ", calculateTotalSalary())
}
