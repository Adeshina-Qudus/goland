package main

func calculatePi(numTerms int) float64 {
	pi := 0.0
	sign := 1.0

	for i := 1; i <= numTerms*2; i += 2 {
		pi += sign * 4 / float64(i)
		sign *= -1
	}
	return pi
}
