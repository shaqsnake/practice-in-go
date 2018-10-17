package primefactors

func factorsOf(n int) []int {
	var factors []int
	if n > 1 {
		return append(factors, 2)
	}
	return factors
}