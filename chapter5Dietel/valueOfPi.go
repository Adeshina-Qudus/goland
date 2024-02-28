package main

func calculatePi(numTerms int) float64 {
	pi := 0.0
	sign := 1.0
	for count := 1; count <= numTerms*2; count += 2 {
		pi += sign * 4 / float64(count)
		sign *= -1
	}
	return pi
}
