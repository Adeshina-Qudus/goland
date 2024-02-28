package main

import (
	"fmt"
	"math"
)

func interest() {
	var principal = 1000.0
	for rate := 5; rate <= 10; rate++ {
		println("rate(", rate, "%)", "\t\tyear\t\tAmount on deposit")
		for year := 1; year <= 10; year++ {
			var amount = principal * math.Pow(1.0+float64(rate), float64(year))
			fmt.Printf("\t\t\t%d\t\t%.1f\n", year, amount)
		}
		println()
		println()
	}

}
