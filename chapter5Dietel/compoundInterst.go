package main

import "math"

func interest() {
	var principal = 1000.0
	for rate := 5; rate <= 10; rate++ {
		println("rate(", rate, "%)", "\t\tyear\t\tAmount on deposit")
		for year := 1; year <= 10; year++ {
			amount := principal * math.Pow(float64(float32(1.0+rate)), float64(year))
			println("\t\t\t", year, "\t\t\t", amount)
		}
		println()
		println()
	}

}
