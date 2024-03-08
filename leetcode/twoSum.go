package main

func twoSum(input []int, target int) []int {
	var result []int
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
