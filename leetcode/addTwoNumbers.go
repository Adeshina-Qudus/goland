package main

import "strconv"

func addTwoNumbers(firstList []int, secondList []int) []int {
	var firstValue string = toString(firstList)
	var secondValue string = toString(secondList)
	intFirstValue, _ := strconv.Atoi(firstValue)
	intSecondValue, _ := strconv.Atoi(secondValue)
	totalOfNumbers := intFirstValue + intSecondValue
	result := strconv.Itoa(totalOfNumbers)
	listOfResult :=

}

func toString(list []int) string {
	var getEachNumber string = ""
	for count := len(list) - 1; count <= 0; count-- {
		getEachNumber += strconv.Itoa(list[count])
	}
	return getEachNumber
}
