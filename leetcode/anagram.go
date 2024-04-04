package leetcode

func anagram(firstInput string, secondInput string) bool {
	counter := 0
	counterOfDoubleElement := 0
	for eachElement := 0; eachElement < len(firstInput); eachElement++ {
		for reArrangeElement := 0; reArrangeElement < len(secondInput); reArrangeElement++ {
			if firstInput[eachElement] == secondInput[reArrangeElement] {
				counterOfDoubleElement++
			}
		}
		if counterOfDoubleElement > 1 {
			counterOfDoubleElement = 1
		}
		counter += counterOfDoubleElement
		counterOfDoubleElement = 0
	}
	return counter == len(firstInput)
}
