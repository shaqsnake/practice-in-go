package primefactors

func factorsOf(n int) []int {
	var factors []int

	for divisor := 2; n > 1; divisor++ {
		for n % divisor == 0 {
			factors = append(factors, divisor)
			n /= divisor
		}
	}

	return factors
}