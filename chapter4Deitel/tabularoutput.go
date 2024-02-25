package main

import (
	"fmt"
)

func tabular() {
	var number = 1
	var n int
	var n2 int
	var n3 int
	var n4 int
	fmt.Println("N\tN2\t N3\t N4\t")
	for count := 1; count <= 5; count++ {
		n = count
		n2 = number * count
		n3 = number * count * count
		n4 = number * count * count * count
		fmt.Println(n, "\t", n2, "\t", n3, "\t", n4, "\t")
		number++
	}
}
