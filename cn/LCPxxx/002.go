func fraction(cont []int) []int {
	l := len(cont)
	a, b := 1, cont[l-1]
	for i := l - 2; i >= 0; i-- {
		a += cont[i] * b
		a, b = b, a
	}
	return []int{b, a}
}
