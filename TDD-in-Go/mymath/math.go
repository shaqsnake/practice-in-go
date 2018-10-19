package mymath

func Add(x ...int) int {
	var result int

	for _, i := range x {
		result += i
	}

	return result
}