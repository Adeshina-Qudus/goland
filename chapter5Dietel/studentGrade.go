package main

import "fmt"

func studentGrade() {
	gradeA := 0
	gradeB := 0
	gradeC := 0
	println("Enter Your Name ")
	var name string
	for count := 1; count <= 5; count++ {
		println("Enter Your Grade : ")
		var input string
		fmt.Scan(&input)
		switch input {
		case "A":
			gradeA++
			break
		case "B":
			gradeB++
			break
		case "C":
			gradeC++
			break
		default:
			println("INVALID GRADE")

		}
	}
	println(name)
	println("THE NUMBER OF STUDENT WHO EARNED GRADE A IS ", gradeA)
	println("THE NUMBER OF STUDENT WHO EARNED GRADE B IS ", gradeB)
	println("THE NUMBER OF STUDENT WHO EARNED GRADE B IS ", gradeC)
}
