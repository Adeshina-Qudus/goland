package main

import "fmt"

func analysis() {
	var passes int
	var failures int
	var studentCounter int
	for studentCounter < 10 {
		println("Enter result (1 = pass, 2 = fail):")
		var studentInput int
		fmt.Scan(&studentInput)
		if studentInput != 1 && studentInput != 2 {
			studentCounter--
		}
		if studentInput == 1 {
			passes++
		}
		if studentInput == 2 {
			failures++
		}
		studentCounter++
	}
	println("Passed: ", passes)
	println("Failures: ", failures)
	if passes > 8 {
		println("Bonus to instructor!")
	}
}
