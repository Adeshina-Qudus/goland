package main

func twoSum(input [4]int, target int) [2]int {
	var result [2]int
	for getIndex := 0; getIndex < len(input); getIndex++ {
		for getElement := 1; getElement < len(input); getElement++ {
			if input[getIndex]+input[getElement] == target {
				result[0] = input[getIndex]
				result[1] = input[getElement]
			}
		}
	}
	return result
}
