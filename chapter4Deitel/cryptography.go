package main

import (
	"fmt"
	"strconv"
)

func encrypt() {
	encryptTo := ""
	fmt.Println("Enter 4 Digit Number To Encrypt  : ")
	var digit string
	fmt.Scan(&digit)
	if len(digit) != 4 {
		for {
			fmt.Println("Enter 4 Digit Number To Encrypt : ")
			fmt.Scan(&digit)
			if len(digit) != 4 {
				continue
			}
		}
	}
	for _, eachNumber := range digit {
		result, _ := strconv.Atoi(string(eachNumber))
		result += 7
		divided := result % 10
		encryptTo += strconv.Itoa(divided)
	}
	println(encryptTo)
}

func decrypt() {
	emptyString := ""
	fmt.Println("Enter 4 Digit Number To Decrypt  : ")
	var digit string
	fmt.Scan(&digit)
	if len(digit) != 4 {
		for {
			fmt.Println("Enter 4 Digit Number To Decrypt : ")
			fmt.Scan(&digit)
			if len(digit) != 4 {
				continue
			}
		}
	}
	for _, eachNumber := range digit {
		for count := 1; count <= 9; count++ {
			result, _ := strconv.Atoi(string(eachNumber))
			countPlusSeven := count + 7
			if countPlusSeven%10 == result {
				emptyString += strconv.Itoa(countPlusSeven - 7)
			}
		}
	}
	println(emptyString)
}

//3452
//0129
