package main

import (
	"fmt"
)

func twoLargest() {
	var firstLargest int
	var secondLargest int
	var temp int
	for counter := 1; counter <= 10; counter++ {
		fmt.Print("Enter a number ")
		var userinput int
		fmt.Scan(&userinput)
		temp = firstLargest
		if userinput > firstLargest {
			firstLargest = userinput
		}
		if firstLargest > temp {
			secondLargest = temp
		}
	}
	println("FirstLargest is :", firstLargest)
	println("SecondLargest is :", secondLargest)
}
