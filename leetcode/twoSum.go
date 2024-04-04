package leetcode

func TwoSum(input []int, target int) []int {
	var result []int
	for getIndex := 0; getIndex < len(input); getIndex++ {
		for getElement := 1; getElement < len(input); getElement++ {
			firstValue := input[getIndex]
			secondValue := input[getElement]
			if firstValue+secondValue == target {
				result = append(result, firstValue)
				result = append(result, secondValue)
			}
		}
	}
	return result
}
