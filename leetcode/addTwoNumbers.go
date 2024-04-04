package leetcode

import (
	"strconv"
)

func addTwoNumbers(firstList []int, secondList []int) []int {
	var firstValue = toString(firstList)
	var secondValue = toString(secondList)
	intFirstValue, _ := strconv.Atoi(firstValue)
	intSecondValue, _ := strconv.Atoi(secondValue)
	totalOfNumbers := intFirstValue + intSecondValue
	result := strconv.Itoa(totalOfNumbers)
	var listOfResult = toList(result)
	return listOfResult
}

func toList(result string) []int {
	var resultList []int
	for eachElement := len(result) - 1; eachElement >= 0; eachElement-- {
		eachvalue, _ := strconv.Atoi(strconv.Itoa(int(result[eachElement])))
		resultList = append(resultList, eachvalue)
	}
	return resultList
}

func toString(list []int) string {
	var getEachNumber = ""
	for count := len(list) - 1; count >= 0; count-- {
		getEachNumber += strconv.Itoa(list[count])
	}
	return getEachNumber
}
