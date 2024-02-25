package main

import (
	"fmt"
)

func salesCommission() {
	var weeklyEarnings float64 = 200
	var totalEarnings float64
	fmt.Print("Enter item sold : ")
	var itemSold float64
	fmt.Scan(&itemSold)
	var ninePercent = (9 * itemSold) / 100
	totalEarnings = ninePercent + weeklyEarnings
	fmt.Println("Total Earnings is ", totalEarnings)
}
