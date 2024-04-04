package main

import (
	"awesomeProject/leetcode"
	"testing"
)

func TestTwoSum(t *testing.T) {
	first := 2
	second := 2
	result := leetcode.SumOfTwoNumbers(first, second)
	if result != 4 {
		t.Error("Expected 4 got ", result)
	}
}
