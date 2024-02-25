package main

func summation() {
	println("Numbers         Summation")
	sum := 0
	for counter := 1; counter <= 100; counter++ {
		sum += counter
		println(counter, " \t\t ", sum)
	}
}
