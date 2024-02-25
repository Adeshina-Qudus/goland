package main

func binaryToDecimal(binaryNumber string) int {
	var decimalResult int
	var power = 1
	for i := len(binaryNumber) - 1; i >= 0; i-- {
		intValue := int(binaryNumber[i] - '0')
		multiply := intValue * power
		decimalResult += multiply
		power *= 2
	}
	return decimalResult
}
