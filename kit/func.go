package kit

func add(a, b int) int {
	return a + b
}
func mul(a, b int) int {
	return a * b
}

func compute(a, b int, operate func(c, d int) int) int {
	c := operate(a, b)
	return c
}
