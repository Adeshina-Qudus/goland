package main

func divisibleBy3() {
	sum := 0
	for count := 1; count <= 30; count++ {
		if count%3 == 0 {
			sum += count
		}
	}
	println("Sum Of Numbers Divisible From To 30 is :", sum)
}
