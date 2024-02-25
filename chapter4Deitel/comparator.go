package main

import "fmt"

func comparator(firstnumber int, secondnumber int) {

	if firstnumber == secondnumber {
		fmt.Print("0")
	}
	if firstnumber > secondnumber {
		fmt.Print("1")
	}
	if firstnumber < secondnumber {
		fmt.Print("-1")
	}

}
