package main

import "fmt"

func main() {
	fmt.Println("Enter the length of the base of a triangle ")
	var input int
	fmt.Scan(&input)
	for firstCounter := 1; firstCounter <= input; firstCounter++ {
		for secondCounter := 1; secondCounter <= firstCounter; secondCounter++ {
			fmt.Print("* ")
		}
		fmt.Println(" ")
	}
}
